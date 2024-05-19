package generate

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var GenerateMigrationCommand = &cli.Command{
	Name:  "migration",
	Usage: "Generate a migration",
	Action: func(c *cli.Context) error {
		fmt.Println("Hello, World! Generating a migration.")
		return nil
	},
}
