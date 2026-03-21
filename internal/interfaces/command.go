package interfaces

import "github.com/urfave/cli/v3"

type Command interface {
	GetCommand() *cli.Command
}
