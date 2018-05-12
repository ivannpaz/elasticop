package template

import (
	"fmt"

	"github.com/urfave/cli"
)

func get(c *cli.Context) error {
	fmt.Println("Get a given template from: ", c.Args().First())
	return nil
}
