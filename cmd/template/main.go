package template

import (
	"fmt"

	"github.com/ivannpaz/elasticop/app/config"
	"github.com/urfave/cli"
)

// New will return the cli.Command definition for all "template" subcommands
func New() cli.Command {
	return cli.Command{
		Name:        "template",
		Description: "Templates management",
		Subcommands: []cli.Command{
			{
				Name:        "list",
				Description: "Get template definition",
				Action: func(c *cli.Context) error {
					cfg := config.Load(c.GlobalString("configuration"))

					if c.NArg() == 0 {
						return fmt.Errorf(
							"Please provide cluster name from the %v",
							cfg.ClusterNames(),
						)
					}

					cluster, err := cfg.Cluster(c.Args().First())
					if err != nil {
						return fmt.Errorf(
							"Requested cluster named %s is not among %v",
							c.Args().First(),
							cfg.ClusterNames(),
						)
					}

					fmt.Printf("List all templates at: %+v\n", cluster)
					return nil
				},
			},
			{
				Name:        "get",
				Description: "Get a given template",
				Action: func(c *cli.Context) error {
					fmt.Println("Get a given template from: ", c.Args().First())
					return nil
				},
			},
			{
				Name:        "put",
				Description: "Overwrite a named template with a new definition",
				Action: func(c *cli.Context) error {
					fmt.Printf("Overwrite template %s with definition %s\n", c.Args().Get(1), c.Args().Get(2))
					return nil
				},
			},
			{
				Name:        "remove",
				Description: "Remove an existing template",
				Action: func(c *cli.Context) error {
					fmt.Println("Removing template: ", c.Args().First())
					return nil
				},
			},
		},
	}
}
