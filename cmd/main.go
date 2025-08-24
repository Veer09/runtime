package main

import (
	"context"
	"log"
	"os"
	"github.com/Veer09/runtime/cmd/operations"
	"github.com/urfave/cli/v3"
)

const specPath = "config.json"

func main(){
	root := "/run/runtime"
	if os.Args[1] == "init" {
		operations.Init()
		return
	}
	cmd := &cli.Command{
		Name: "runtime",
		Usage: "runtime",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "root",
				Value: root,
			},
		},
		Commands: []*cli.Command{
			CreateCmd,
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}