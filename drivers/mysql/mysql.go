package mysql

import (
	"database/sql"

	"github.com/ecorreiax/flare/internal/query"
)

type MysqlDriver struct{}

func (d *MysqlDriver) Open() (*sql.DB, error) {
	return query.SQLOpen(d)
}

func (d *MysqlDriver) Create() error {
	return query.SQLCreate(d)
}

func (d *MysqlDriver) Drop() error {
	return query.SQLDrop(d)
}
