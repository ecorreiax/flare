package flare

import (
	"database/sql"
	"errors"
)

type Driver interface {
	Open() (*sql.DB, error)
	Create() error
	Drop() error
}

var drivers = map[string]Driver{}

func (db *Database) Driver() (Driver, error) {
	drv := drivers[db.URL.Scheme]

	if drv == nil {
		return nil, errors.New("invalid driver")
	}

	return drv, nil
}

func RegisterDriver(driver Driver, scheme string) {
	drivers[scheme] = driver
}
