package config

import (
	"errors"
	"fmt"
	"log"
	"net/url"

	"github.com/pelletier/go-toml"
)

type Config struct {
	Driver   string
	User     string
	Name     string
	Host     string
	Port     string
	Password string
}

var AppConfig Config

func init() {
	tree, err := toml.LoadFile("flare.toml")
	if err != nil {
		log.Fatal(err)
	}

	AppConfig = Config{
		Driver:   tree.Get("Driver").(string),
		User:     tree.Get("User").(string),
		Name:     tree.Get("Name").(string),
		Host:     tree.Get("Host").(string),
		Port:     tree.Get("Port").(string),
		Password: tree.Get("Password").(string),
	}

	if err := validateConfig(AppConfig); err != nil {
		log.Fatal(err)
	}
}

func validateConfig(c Config) error {
	if len(c.Driver) == 0 {
		return errors.New("database driver key missing")
	}
	if len(c.User) == 0 {
		return errors.New("database user key missing")
	}
	if len(c.Password) == 0 {
		return errors.New("database password key missing")
	}
	if len(c.Host) == 0 {
		return errors.New("database host key missing")
	}
	if len(c.Port) == 0 {
		return errors.New("database port key missing")
	}

	return nil
}

func (c *Config) BuildConnStr() (*url.URL, error) {
	var connStr string

	switch c.Driver {
	case "postgres":
		connStr = fmt.Sprintf(
			"postgres://%s:%s@%s:%s/postgres?sslmode=disable",
			c.User,
			c.Password,
			c.Host,
			c.Port,
		)
	case "mysql":
		connStr = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/",
			c.User,
			c.Password,
			c.Host,
			c.Port,
		)
	default:
		return nil, errors.New("database provider missing or invalid")
	}

	return url.Parse(connStr)
}
