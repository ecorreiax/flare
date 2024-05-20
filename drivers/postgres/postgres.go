package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/ecorreiax/flare/internal/config"
	"github.com/ecorreiax/flare/internal/flare"
	_ "github.com/lib/pq"
)

type PostgresDriver struct{}

func (d *PostgresDriver) Open() (*sql.DB, error) {
	url, err := config.AppConfig.BuildConnStr()
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return sql.Open("postgres", url.String())
}

func (d *PostgresDriver) Create() error {
	dbName := config.AppConfig.Name

	db, err := d.Open()
	if err != nil {
		return err
	}
	defer flare.MustClose(db)

	_, err = db.Exec("create database " + dbName)

	return err
}

func (d *PostgresDriver) Drop() error {
	dbName := config.AppConfig.Name

	db, err := d.Open()
	if err != nil {
		return err
	}
	defer flare.MustClose(db)

	_, err = db.Exec(fmt.Sprintf("drop database if exists %s", dbName))

	return err
}
