package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/ecorreiax/flare/internal/config"
	"github.com/urfave/cli/v2"
)

func GenerateMigrationFile(c *cli.Context) error {
	name := c.Args().Get(0)
	if len(name) == 0 {
		return errors.New("migration name wasn't provided")
	}

	path := fmt.Sprintf("%s/migrate", config.AppConfig.Path)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}

	filename := fmt.Sprintf("%s_%s.sql", time.Now().Format("20060102150405"), name)
	filepath := filepath.Join(config.AppConfig.Path, "migrate", filename)

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}
