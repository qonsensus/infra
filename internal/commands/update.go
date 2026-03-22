package commands

import (
	commandrunners "github.com/qonsensus/infra/internal/commandRunners"
	"github.com/qonsensus/infra/internal/interfaces"
	"github.com/urfave/cli/v3"
)

type UpdateCommand struct{}

func (c *UpdateCommand) GetCommand() *cli.Command {
	return &cli.Command{
		Name:  "update",
		Usage: "Recreate the docker-compose.yaml file. This MUST be run after updating the config.toml file.",
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:      "target",
				UsageText: "<target-dir>",
				Value:     ".",
			},
		},
		Action: (&commandrunners.UpdateCommandRunner{}).Run,
	}
}

var _ interfaces.Command = (*UpdateCommand)(nil)
