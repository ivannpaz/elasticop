package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/ivannpaz/elasticop/app/config"
	"github.com/urfave/cli"
)

var version = "unset"

func main() {
	app := cli.NewApp()
	app.Name = "Elastic Operator"
	app.Usage = ""
	app.Description = "Perform common tasks while working with Elasticsearch in a very simple manner"
	app.Version = version

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "configuration, c",
			Value: "./config.yml",
			Usage: "Configuration file",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "config",
			Usage: "List configuration values",
			Action: func(c *cli.Context) error {
				cfg := config.Load(c.GlobalString("configuration"))

				if json, err := json.Marshal(cfg); err == nil {
					fmt.Printf("%s", json)
				}

				return nil
			},
		},
		{
			Name:  "template",
			Usage: "Template management",
			Subcommands: []cli.Command{
				{
					Name:  "get",
					Usage: "Get template definition",
					Action: func(c *cli.Context) error {
						fmt.Println("new task template: ", c.Args().First())
						return nil
					},
				},
				{
					Name:  "add",
					Usage: "add a new template",
					Action: func(c *cli.Context) error {
						fmt.Println("new task template: ", c.Args().First())
						return nil
					},
				},
				{
					Name:  "remove",
					Usage: "remove an existing template",
					Action: func(c *cli.Context) error {
						fmt.Println("removed task template: ", c.Args().First())
						return nil
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
