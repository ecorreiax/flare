package database

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var DatabaseMigrateCommand = &cli.Command{
	Name:  "migrate",
	Usage: "Migrate the database",
	Action: func(c *cli.Context) error {
		fmt.Println("Hello, World! Migrating the database.")
		// flare g migration user
		// - creates an empty user migration with minimal template
		//
		// flare g migration user name email password is_admin:bool
		// - creates an user migration with columns
		return nil
	},
}
