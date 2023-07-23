package storage_fx

import (
	"fmt"
	"members/config"
	"members/storage"
	"members/storage/mem"
	"members/storage/mysql"
	pgres "members/storage/pgx"
	"members/utils"
)

func create_sql(prov config.ConfigProvider, create storage.Initializer) (storage.Store, error) {
	if db, err := create(prov); err != nil {
		return nil, err
	} else {
		if prov.GetConfig().Storage.Debug {
			db.AddQueryHook(&qh{})
		}
		return storage.NewSql(db), nil
	}
}

func New(prov config.ConfigProvider) (storage.Store, error) {
	st := prov.GetConfig().Storage.Kind

	switch {
	case utils.AnyEq([]string{
		"mem", "memory", "sqlite",
	}, st) || utils.ZeroStr(st):
		return create_sql(
			prov,
			mem.New,
		)
	case utils.AnyEq([]string{
		"postgres", "pg",
	}, st):
		return create_sql(
			prov,
			pgres.New,
		)
	case utils.AnyEq([]string{
		"mysql", "maria",
	}, st):
		return create_sql(
			prov,
			mysql.New,
		)
	}
	return nil, fmt.Errorf("invalid storage type: %s", st)
}

func with_create(config config.ConfigProvider) {
	config.GetConfig().Storage.Create = true
}
