package template

import (
	"fmt"

	"github.com/urfave/cli"
)

func put(c *cli.Context) error {
	fmt.Printf("Overwrite template %s with definition %s\n", c.Args().Get(1), c.Args().Get(2))
	return nil
}
