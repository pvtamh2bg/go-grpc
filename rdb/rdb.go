package rdb

import (
	"gorm.io/gorm"
)

type Connection struct {
	*gorm.DB
}

type Env struct {
	Host     string // EMZ_RDB_HOST
	User     string // EMZ_RDB_USER
	Password string // EMZ_RDB_PASSWORD
	Dbname   string // EMZ_RDB_DBNAME
	Driver   string // EMZ_RDB_DRIVER
}

func Connect() (con Connection, err error) {
	return con, nil
}
