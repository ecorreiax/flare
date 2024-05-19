package main

import (
	"os"

	"github.com/urfave/cli/v2"

	"github.com/ecorreiax/flare/cmd/database"
	"github.com/ecorreiax/flare/cmd/generate"
)

func main() {
	app := &cli.App{
		Name:  "flare",
		Usage: "A lightweight migration tool for Go",
		Commands: []*cli.Command{
			{
				Name:    "database",
				Aliases: []string{"db"},
				Usage:   "Database related commands",
				Subcommands: []*cli.Command{
					database.DatabaseCreateCommand,
					database.DatabaseDropCommand,
				},
			},
			{
				Name:    "generate",
				Aliases: []string{"g"},
				Usage:   "Generate related commands",
				Subcommands: []*cli.Command{
					generate.GenerateMigrationCommand,
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		println(err.Error())
	}
}
