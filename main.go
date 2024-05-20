package main

import (
	"os"

	"github.com/urfave/cli/v2"

	"github.com/ecorreiax/flare/cmd"
	"github.com/ecorreiax/flare/drivers/mysql"
	"github.com/ecorreiax/flare/drivers/postgres"
	"github.com/ecorreiax/flare/internal/config"
	"github.com/ecorreiax/flare/internal/flare"
)

func init() {
	flare.RegisterDriver(&mysql.MysqlDriver{}, "mysql")
	flare.RegisterDriver(&postgres.PostgresDriver{}, "postgres")
}

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
					{
						Name:  "create",
						Usage: "Create the database",
						Action: wrap(func(db *flare.Database, _ *cli.Context) error {
							return db.Create()
						}),
					},
					{
						Name:  "drop",
						Usage: "Drop the database",
						Action: wrap(func(db *flare.Database, _ *cli.Context) error {
							return db.Drop()
						}),
					},
				},
			},
			{
				Name:    "generate",
				Aliases: []string{"g"},
				Usage:   "Generate templated files",
				Subcommands: []*cli.Command{
					{
						Name:  "migration",
						Usage: "Create a new migration",
						Action: func(c *cli.Context) error {
							return cmd.GenerateMigrationFile(c)
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		println(err.Error())
	}
}

func wrap(f func(*flare.Database, *cli.Context) error) cli.ActionFunc {
	return func(c *cli.Context) error {
		url, err := config.AppConfig.BuildConnStr()
		if err != nil {
			return err
		}
		db := flare.New(url)
		return f(db, c)
	}
}
