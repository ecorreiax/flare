package flare

import (
	"io"
	"net/url"
)

type Database struct {
	URL *url.URL
}

func (db *Database) Create() error {
	drv, err := db.Driver()
	if err != nil {
		return err
	}

	return drv.Create()
}

func (db *Database) Drop() error {
	drv, err := db.Driver()
	if err != nil {
		return err
	}

	return drv.Drop()
}

func New(url *url.URL) *Database {
	return &Database{
		URL: url,
	}
}

func MustClose(c io.Closer) {
	if err := c.Close(); err != nil {
		panic(err)
	}
}
