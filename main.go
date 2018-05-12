package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/ivannpaz/elasticop/app/config"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	config := config.Load("config.yml")

	app.Commands = []cli.Command{
		{
			Name:  "config",
			Usage: "list conguration",
			Action: func(c *cli.Context) error {
				if json, err := json.Marshal(config); err == nil {
					fmt.Print(json)
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
