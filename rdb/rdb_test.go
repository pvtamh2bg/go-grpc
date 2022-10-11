package rdb_test

import (
	"context"
	"fmt"
	"grpc/test/rdb"
	"log"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	rdbLogger "gorm.io/gorm/logger"
)

const searchPath = "auth"

var JST = time.FixedZone("JST", 9*60*60)

type testEnv struct {
	Host     string // EMZ_RDB_HOST
	User     string // EMZ_RDB_USER
	Password string // EMZ_RDB_PASSWORD
	Port     string // EMZ_RDB_PORT
	Dbname   string // EMZ_RDB_DBNAME
	Driver   string // EMZ_RDB_DRIVER
}

func ConnectForTest() (conn rdb.Connection, err error) {
	// load env
	err = godotenv.Load("../test.env")
	if err != nil {
		log.Printf("error loading env file err=%+v", err) // TODO 問題点 No.46
	}
	var rdbEnv testEnv
	_ = envconfig.Process("EMZ_RDB", &rdbEnv)
	// connect to RDB
	dsn := fmt.Sprintf("host=%s user=%s "+
		"password=%s dbname=%s port=%s search_path=%s sslmode=disable",
		rdbEnv.Host, rdbEnv.User, rdbEnv.Password, rdbEnv.Dbname, rdbEnv.Port, searchPath)

	conn.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	// log sql executing
	conn.Logger = conn.DB.Logger.LogMode(rdbLogger.Info)

	return
}

func initDBForTests(ctx context.Context, t *testing.T, conn rdb.Connection) {
	require.NoError(t, conn.WithContext(ctx).Exec(`TRUNCATE TABLE auth.users RESTART IDENTITY;`).Error)
	require.NoError(t, conn.WithContext(ctx).Exec(`TRUNCATE TABLE auth.post_invitees RESTART IDENTITY;`).Error)
}
