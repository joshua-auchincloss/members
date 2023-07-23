package mysql

import (
	"database/sql"
	"fmt"
	"members/config"
	"members/utils"

	"github.com/rs/zerolog/log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

var (
	mysqlhost        = "localhost"
	mysqlport uint32 = 4900
	mysqldb          = "mysql"
)

func New(prov config.ConfigProvider) (*bun.DB, error) {
	cfg := prov.GetConfig()
	host := cfg.Storage.URI
	if utils.ZeroStr(host) {
		host = mysqlhost
	}
	port := cfg.Storage.Port
	if utils.IsZero(port) {
		port = mysqlport
	}
	dbname := cfg.Storage.DB
	if utils.ZeroStr(dbname) {
		dbname = mysqldb
	}
	var sslstr string
	if cfg.Storage.SSL {
		sslstr = "?sslmode=enabled"
	}
	connstr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s%s",
		cfg.Storage.Username,
		cfg.Storage.Password,
		host,
		port,
		dbname,
		sslstr,
	)
	log.Print(connstr)
	sqldb, err := sql.Open("mysql", connstr)
	if err != nil {
		return nil, err
	}
	db := bun.NewDB(sqldb, mysqldialect.New())
	return db, nil
}
