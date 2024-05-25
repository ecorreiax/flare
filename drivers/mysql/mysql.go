package mysql

import (
	"database/sql"

	"github.com/ecorreiax/flare/internal/query"
)

type MysqlDriver struct{}

const DBScheme = "mysql"

func (d *MysqlDriver) Open() (*sql.DB, error) {
	return query.SQLOpen(d, DBScheme)
}

func (d *MysqlDriver) Create() error {
	return query.SQLCreate(d)
}

func (d *MysqlDriver) Drop() error {
	return query.SQLDrop(d)
}
