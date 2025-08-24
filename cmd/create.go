package main

import (
	"context"
	"errors"

	"github.com/Veer09/runtime/cmd/operations"
	"github.com/urfave/cli/v3"
)

var CreateCmd = &cli.Command{
	Name: "create",
	Usage: "create a new container",
	Action: func(ctx context.Context, cmd *cli.Command) error {
		id := cmd.Args().First()
		bundle := cmd.Args().Get(1)
		if id == "" {
			return errors.New("container id is required")
		}
		if bundle == "" {
			return errors.New("bundle is required")
		}
		err := operations.Create(cmd, id, bundle)
		if err != nil {
			return err
		}
		
		return nil
	},

}