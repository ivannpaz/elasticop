package main

import (
	"log"
	"os"

	"github.com/ivannpaz/elasticop/cmd/configuration"
	"github.com/ivannpaz/elasticop/cmd/template"

	"github.com/urfave/cli"
)

var version = "unset"

func main() {
	app := cli.NewApp()
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
		configuration.New(),
		template.New(),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
