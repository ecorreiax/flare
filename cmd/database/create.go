package database

import (
	"fmt"

	"github.com/ecorreiax/flare/drivers/mysql"
	"github.com/ecorreiax/flare/drivers/postgres"
	"github.com/ecorreiax/flare/internal/config"
	"github.com/urfave/cli/v2"
)

var DatabaseCreateCommand = &cli.Command{
	Name:  "create",
	Usage: "Create the database",
	Action: func(c *cli.Context) error {
		switch config.AppConfig.Driver {
		case "postgres":
			return postgres.CreateDatabase()
		case "mysql":
			return mysql.CreateDatabase()
		default:
			fmt.Println("database provider missing or invalid")
		}

		return nil
	},
}
