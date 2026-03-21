package commands

import (
	commandrunners "github.com/qonsensus/infra/internal/commandRunners"
	"github.com/qonsensus/infra/internal/interfaces"
	"github.com/urfave/cli/v3"
)

type InitCommand struct{}

func (c *InitCommand) GetCommand() *cli.Command {
	return &cli.Command{
		Name:  "init",
		Usage: "Initialize Quonsensus",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "output",
				Usage:   "Output location for the initialized files",
				Value:   "./quonsensus",
				Aliases: []string{"o"},
			},
		},
		Action: (&commandrunners.InitCommandRunner{}).Run,
	}
}

var _ interfaces.Command = (*InitCommand)(nil)
