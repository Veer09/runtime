package operations

import (
	"fmt"
	"os"

	"github.com/Veer09/runtime/internal"
	"github.com/urfave/cli/v3"
)

func Create(cmd *cli.Command, id string, bundle string) error {
	if err := os.Chdir(bundle); err != nil {
		return err
	}
	spec, err := readSpec()
	if err != nil {
		return err
	}
	root := cmd.String("root")
	fmt.Println(root)
	err = runtime.Create(root, id, spec)
	if err != nil {
		return err
	}
	return nil
}

