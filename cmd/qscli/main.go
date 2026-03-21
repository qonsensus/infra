package main

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "qscli",
		Usage: "A sample CLI using urfave/cli/v3",
		Action: func(ctx context.Context, c *cli.Command) error {
			fmt.Println("Hello from urfave/cli/v3!")
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
