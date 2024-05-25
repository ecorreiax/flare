package query

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/ecorreiax/flare/internal/config"
	"github.com/ecorreiax/flare/internal/flare"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func SQLOpen(d flare.Driver, DBScheme string) (*sql.DB, error) {
	url, err := config.AppConfig.BuildConnStr()
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return sql.Open(DBScheme, url.String())
}

func SQLCreate(d flare.Driver) error {
	dbName := config.AppConfig.Name

	db, err := d.Open()
	if err != nil {
		return err
	}
	defer flare.MustClose(db)

	_, err = db.Exec("create database " + dbName)

	return err
}

func SQLDrop(d flare.Driver) error {
	dbName := config.AppConfig.Name

	db, err := d.Open()
	if err != nil {
		return err
	}
	defer flare.MustClose(db)

	_, err = db.Exec(fmt.Sprintf("drop database if exists %s", dbName))

	return err
}
