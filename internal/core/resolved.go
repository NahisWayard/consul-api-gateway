package core

type ResolvedService struct {
	ConsulNamespace string
	Service         string
}

type ResolvedRouteType int

const (
	ResolvedHTTPRouteType ResolvedRouteType = iota
	ResolvedTCPRouteType
	ResolvedTLSRouteType
	ResolvedUDPRouteType
)

type ResolvedRoute interface {
	Type() ResolvedRouteType
	Meta() map[string]string
	Name() string
	Namespace() string
}

type ResolvedListener struct {
	Name         string
	Hostname     string
	Port         int
	Protocol     string
	Certificates []string
	Routes       []ResolvedRoute
}

type GatewayID struct {
	ConsulNamespace string
	Service         string
}

type ResolvedGateway struct {
	ID        GatewayID
	Meta      map[string]string
	Listeners []ResolvedListener
}