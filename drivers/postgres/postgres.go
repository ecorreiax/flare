package postgres

import (
	"database/sql"

	"github.com/ecorreiax/flare/internal/query"
)

type PostgresDriver struct{}

const DBScheme = "postgres"

func (d *PostgresDriver) Open() (*sql.DB, error) {
	return query.SQLOpen(d, DBScheme)
}

func (d *PostgresDriver) Create() error {
	return query.SQLCreate(d)
}

func (d *PostgresDriver) Drop() error {
	return query.SQLDrop(d)
}
