package commandrunners

import (
	"context"

	"github.com/qonsensus/infra/internal/interfaces"
	"github.com/qonsensus/infra/internal/models"
	"github.com/urfave/cli/v3"
)

type InitCommandRunner struct{}

func (r *InitCommandRunner) Run(ctx context.Context, cmd *cli.Command) error {
	basePath := cmd.String("output")
	version := cmd.String("version")

	config := models.NewConfigToml(version, basePath)
	if err := config.SaveToFile(); err != nil {
		return err
	}
	return nil
}

var _ interfaces.CommandRunner = (*InitCommandRunner)(nil)
