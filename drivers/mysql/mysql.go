package mysql

import (
	"database/sql"
	"fmt"

	"github.com/ecorreiax/flare/internal/config"
	"github.com/ecorreiax/flare/internal/database"
)

func OpenConnection() (*sql.DB, error) {
	return sql.Open("mysql", config.AppConfig.BuildConnStr())
}

func CreateDatabase() error {
	dbName := config.AppConfig.Name

	db, err := OpenConnection()
	if err != nil {
		return err
	}
	defer database.MustClose(db)

	_, err = db.Exec("CREATE DATABASE " + dbName)

	return err
}

func DropDatabase() error {
	dbName := config.AppConfig.Name

	db, err := OpenConnection()
	if err != nil {
		return err
	}
	defer database.MustClose(db)

	_, err = db.Exec(fmt.Sprintf("drop database if exists %s with (force)", dbName))

	return err
}
