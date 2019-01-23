package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

var (
	version = "0.0.0"
	build   = "0"
)

func main() {
	app := cli.NewApp()
	app.Name = "azure storage plugin"
	app.Usage = "azure storage plugin"
	app.Action = run
	app.Version = fmt.Sprintf("%s+%s", version, build)
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "account-key",
			Usage:  "account key",
			EnvVar: "PLUGIN_ACCOUNT_KEY,AZURE_STORAGE_ACCOUNT_KEY",
		},
		cli.StringFlag{
			Name:   "account",
			Usage:  "account",
			EnvVar: "PLUGIN_ACCOUNT,PLUGIN_STORAGE_ACCOUNT,AZURE_STORAGE_ACCOUNT",
		},
		cli.StringFlag{
			Name:   "container",
			Usage:  "container",
			EnvVar: "PLUGIN_CONTAINER,AZURE_STORAGE_CONTAINER",
		},
		cli.StringFlag{
			Name:   "source",
			Usage:  "source",
			EnvVar: "PLUGIN_SOURCE,AZURE_STORAGE_SOURCE",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{
		Config: Config{
			AccountKey: c.String("account-key"),
			Account:    c.String("account"),
			Container:  c.String("container"),
			Source:     c.String("source"),
		},
	}

	if plugin.Config.AccountKey == "" {
		return errors.New("Missing account_key")
	}

	if plugin.Config.Container == "" {
		return errors.New("Missing container")
	}

	return plugin.Exec()
}
