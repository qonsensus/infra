package commands

import (
	"github.com/qonsensus/infra/internal"
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
			&cli.StringFlag{
				Name:    "version",
				Usage:   "Version of Quonsensus to initialize",
				Value:   internal.CurrentVersion,
				Aliases: []string{"v"},
			},
		},
		Action: (&commandrunners.InitCommandRunner{}).Run,
	}
}

var _ interfaces.Command = (*InitCommand)(nil)
