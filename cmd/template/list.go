package template

import (
	"fmt"

	"github.com/ivannpaz/elasticop/app/config"
	"github.com/urfave/cli"
)

func list(c *cli.Context) error {
	cfg := config.Load(c.GlobalString("configuration"))

	if c.NArg() == 0 {
		return fmt.Errorf(
			"Please provide cluster name from %+v",
			cfg.ClusterNames(),
		)
	}

	cluster, err := cfg.Cluster(c.Args().First())
	if err != nil {
		return fmt.Errorf(
			"Requested cluster named %s is not configured. Available names are %v",
			c.Args().First(),
			cfg.ClusterNames(),
		)
	}

	fmt.Printf("List all templates at: %+v\n", cluster)
	return nil
}
