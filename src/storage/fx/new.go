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

func create_sql(prov config.ConfigProvider, create storage.Initializer, kind storage.StorageType) (storage.Store, error) {
	if db, err := create(prov); err != nil {
		return nil, err
	} else {
		if prov.GetConfig().Storage.Debug {
			db.AddQueryHook(&qh{})
		}
		return storage.NewSql(db, kind), nil
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
			storage.Memory,
		)
	case utils.AnyEq([]string{
		"postgres", "pg",
	}, st):
		return create_sql(
			prov,
			pgres.New,
			storage.Postgres,
		)
	case utils.AnyEq([]string{
		"mysql", "maria",
	}, st):
		return create_sql(
			prov,
			mysql.New,
			storage.Mysql,
		)
	}
	return nil, fmt.Errorf("invalid storage type: %s", st)
}
