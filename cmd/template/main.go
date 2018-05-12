package template

import (
	"github.com/urfave/cli"
)

// New will return the cli.Command definition for all "template" subcommands
func New() cli.Command {
	return cli.Command{
		Name:        "template",
		Description: "Templates management",
		Action: func(c *cli.Context) error {
			return list(c)
		},
		Subcommands: []cli.Command{
			{
				Name:        "get",
				Description: "Get a given template",
				Action: func(c *cli.Context) error {
					return get(c)
				},
			},
			{
				Name:        "put",
				Description: "Overwrite a named template with a new definition",
				Action: func(c *cli.Context) error {
					return put(c)
				},
			},
			{
				Name:        "remove",
				Description: "Remove an existing template",
				Action: func(c *cli.Context) error {
					return remove(c)
				},
			},
		},
	}
}
