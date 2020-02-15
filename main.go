package main

import (
	"errors"
	"log"
	"os"

	"github.com/urfave/cli"
)

var (
	version = "unknown"
)

func main() {
	app := cli.NewApp()
	app.Name = "azure storage plugin"
	app.Usage = "azure storage plugin"
	app.Action = run
	app.Version = version
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
		cli.StringFlag{
			Name:   "destination",
			Usage:  "destination",
			EnvVar: "PLUGIN_DESTINATION,AZURE_STORAGE_DESTINATION",
		},
		cli.StringFlag{
			Name:   "operation",
			Usage:  "operation",
			EnvVar: "PLUGIN_OPERATION,AZURE_STORAGE_OPERATION",
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
			Destination: c.String("destination"),
			Operation:  c.String("operation"),
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
