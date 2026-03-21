package main

import (
	"context"
	"fmt"
	"os"

	"github.com/qonsensus/infra/internal/commands"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "qscli",
		Usage: "A sample CLI using urfave/cli/v3",
		Commands: []*cli.Command{
			(&commands.InitCommand{}).GetCommand(),
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
