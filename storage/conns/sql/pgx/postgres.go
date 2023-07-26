package pgx

import (
	"fmt"
	"members/config"
	"members/utils"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

var (
	pghost           = "localhost"
	pgport    uint32 = 5432
	defaultdb        = "postgres"
)

func New(prov config.ConfigProvider) (*bun.DB, error) {
	cfg := prov.GetConfig()
	host := cfg.Storage.URI
	if utils.ZeroStr(host) {
		host = pghost
	}
	port := cfg.Storage.Port
	if utils.IsZero(port) {
		port = pgport
	}
	dbname := cfg.Storage.DB
	if utils.ZeroStr(dbname) {
		dbname = defaultdb
	}
	var sslstr string
	if cfg.Storage.SSL {
		sslstr = "?sslmode=enabled"
	}
	connstr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s%s",
		cfg.Storage.Username,
		cfg.Storage.Password,
		host,
		port,
		dbname,
		sslstr,
	)
	config, err := pgx.ParseConfig(connstr)
	if err != nil {
		return nil, err
	}
	sqldb := stdlib.OpenDB(*config)
	db := bun.NewDB(sqldb, pgdialect.New())
	return db, nil
}
