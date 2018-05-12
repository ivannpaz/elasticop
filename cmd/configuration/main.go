package configuration

import (
	"encoding/json"
	"fmt"

	"github.com/ivannpaz/elasticop/app/config"
	"github.com/urfave/cli"
)

// New returns a configured urfave/cli command with any subcommands it may have
func New() cli.Command {
	return cli.Command{
		Name:        "configuration",
		Description: "List configuration values",
		Action: func(c *cli.Context) error {
			cfg := config.Load(c.GlobalString("configuration"))

			if json, err := json.Marshal(cfg); err == nil {
				fmt.Printf("%s", json)
			}

			return nil
		},
	}
}
