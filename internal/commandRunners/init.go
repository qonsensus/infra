package commandrunners

import (
	"context"
	"os"
	"path/filepath"

	"github.com/compose-spec/compose-go/v2/types"
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

	compose := &types.Project{
		Name: "quonsensus",
		Services: types.Services{
			"postgres": models.NewPostgresService(config).ServiceConfig,
		},
	}

	if err := r.SaveComposeFile(compose, basePath); err != nil {
		return err
	}

	return nil
}

func (r *InitCommandRunner) SaveComposeFile(compose *types.Project, basePath string) error {
	composeString, err := compose.MarshalYAML()
	if err != nil {
		return err
	}
	finalPath := filepath.Join(basePath, "docker-compose.yaml")
	return os.WriteFile(finalPath, []byte(composeString), 0644)
}

var _ interfaces.CommandRunner = (*InitCommandRunner)(nil)
