package interfaces

import (
	"context"

	"github.com/urfave/cli/v3"
)

type CommandRunner interface {
	Run(ctx context.Context, cmd *cli.Command) error
}
