package tcproutes

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/hashicorp/consul-api-gateway/internal/common"
	"github.com/mitchellh/cli"
)

type GetCommand struct {
	*common.ClientCLIWithNamespace
}

func NewGetCommand(ctx context.Context, ui cli.Ui, logOutput io.Writer) *GetCommand {
	return &GetCommand{
		ClientCLIWithNamespace: common.NewClientCLIWithNamespace(ctx, getHelp, getSynopsis, ui, logOutput, "get"),
	}
}

func (c *GetCommand) Run(args []string) int {
	if err := c.Parse(args); err != nil {
		return c.Error("parsing command line flags", err)
	}

	name := c.Flags.Arg(0)
	if name == "" {
		return c.Error("parsing arguments", errors.New("a name parameter must be provided"))
	}

	client, err := c.CreateClient()
	if err != nil {
		return c.Error("creating the client", err)
	}

	route, err := client.V1().GetTCPRouteInNamespace(c.Context(), c.Namespace(), name)
	if err != nil {
		return c.Error("sending the request", err)
	}

	return c.Success(fmt.Sprintf("Successfully retrieved http-route: %v", route))
}

const (
	getSynopsis = "Gets a TCPRoute"
	getHelp     = `
Usage: consul-api-gateway tcp-routes get [options] NAME

  Gets a TCPRoute with the given NAME.

  Additional flags and more advanced use cases are detailed below.
`
)
