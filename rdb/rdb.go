package rdb

import (
	"fmt"
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	rdbLogger "gorm.io/gorm/logger"
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

const searchPath = "auth"

var JST = time.FixedZone("JST", 9*60*60)

func Connect() (conn Connection, err error) {
	// rdb env
	var rdbEnv Env
	_ = envconfig.Process("EMZ_RDB", &rdbEnv)

	// dns
	dsn := fmt.Sprintf("host=%s user=%s "+
		"password=%s dbname=%s search_path=%s sslmode=disable",
		rdbEnv.Host, rdbEnv.User, rdbEnv.Password, rdbEnv.Dbname, searchPath)

	// connect to db
	conn.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("connect db fail: %v", err)
	}

	// log sql executing
	conn.Logger = conn.DB.Logger.LogMode(rdbLogger.Info)

	return
}
