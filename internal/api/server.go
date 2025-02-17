package api

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/hashicorp/consul-api-gateway/internal/api/apiinternal"
	v1 "github.com/hashicorp/consul-api-gateway/internal/api/v1"
	consul "github.com/hashicorp/consul/api"
	"github.com/hashicorp/go-hclog"
)

type ServerConfig struct {
	Logger          hclog.Logger
	Consul          *consul.Client
	Address         string
	CertFile        string
	KeyFile         string
	ShutdownTimeout time.Duration

	Validator v1.Validator
	Name      string
	Namespace string

	// info for bootstrapping our deployments
	Bootstrap apiinternal.BootstrapConfiguration
}

type Server struct {
	logger          hclog.Logger
	server          *http.Server
	certFile        string
	keyFile         string
	shutdownTimeout time.Duration
}

func NewServer(config ServerConfig) *Server {
	router := chi.NewRouter()
	router.Mount("/api/v1", v1.NewServer("/api/v1", config.Validator, config.Name, config.Namespace, config.Consul, config.Logger))
	router.Mount("/api/internal", apiinternal.NewServer("/api/internal", config.Bootstrap, config.Consul, config.Logger))

	return &Server{
		logger: config.Logger,
		server: &http.Server{
			Handler: router,
			Addr:    config.Address,
		},
		certFile:        config.CertFile,
		keyFile:         config.KeyFile,
		shutdownTimeout: config.ShutdownTimeout,
	}
}

// Run starts the API server
func (s *Server) Run(ctx context.Context) error {
	errs := make(chan error, 1)
	go func() {
		if s.certFile != "" && s.keyFile != "" {
			s.logger.Info("Certificate file and private key file provided, serving API over HTTPS", "address", s.server.Addr)
			errs <- s.server.ListenAndServeTLS(s.certFile, s.keyFile)
		} else {
			s.logger.Info("TLS certificate configuration not provided, serving API over HTTP", "address", s.server.Addr)
			errs <- s.server.ListenAndServe()
		}
	}()

	for {
		select {
		case err := <-errs:
			return err
		case <-ctx.Done():
			return s.Shutdown()
		}
	}
}

// Shutdown attempts to gracefully shutdown the server, it
// is called automatically when the context passed into the
// Run function is canceled.
func (s *Server) Shutdown() error {
	if s.server != nil {
		ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
		defer cancel()

		return s.server.Shutdown(ctx)
	}
	return nil
}
