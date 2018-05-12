package template

import (
	"fmt"

	"github.com/urfave/cli"
)

func remove(c *cli.Context) error {
	fmt.Println("Removing template: ", c.Args().First())
	return nil
}
