package commandrunners

import (
	"context"
	"fmt"

	"github.com/qonsensus/infra/internal/interfaces"
	"github.com/urfave/cli/v3"
)

type InitCommandRunner struct{}

func (r *InitCommandRunner) Run(ctx context.Context, cmd *cli.Command) error {
	fmt.Println("Running InitCommandRunner...")
	return nil
}

var _ interfaces.CommandRunner = (*InitCommandRunner)(nil)
