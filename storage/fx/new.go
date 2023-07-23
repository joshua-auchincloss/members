package storage_fx

import (
	"fmt"
	"members/config"
	"members/logging"
	"members/storage"
	"members/storage/mem"
	"members/storage/mysql"
	pgres "members/storage/pgx"
	"members/utils"

	"github.com/rs/zerolog"
)

func create_sql(prov config.ConfigProvider, root *zerolog.Logger, create storage.Initializer) (storage.Store, error) {
	store := prov.GetConfig().Storage
	sub := logging.WithSub(root, "storage", func(ctx zerolog.Context) zerolog.Context {
		ctx = ctx.Str("storage-kind", store.Kind).
			Str("storage-uri", store.URI).
			Bool("storage-debug", store.Debug)
		if store.Debug {
			ctx = ctx.
				Uint32("storage-port", store.Port).
				Bool("storage-create-on-init", store.Create).
				Bool("storage-drop-on-init", store.Drop)
		}
		return ctx
	})
	if db, err := create(prov); err != nil {
		return nil, err
	} else {
		sql := storage.NewSql(db)
		sql.WithLogger(sub)
		if prov.GetConfig().Storage.Debug {
			db.AddQueryHook(&qh{sub})
		}
		return sql, nil
	}
}

func New(prov config.ConfigProvider, root *zerolog.Logger) (storage.Store, error) {
	st := prov.GetConfig().Storage.Kind
	switch {
	case utils.AnyEq([]string{
		"mem", "memory", "sqlite",
	}, st) || utils.ZeroStr(st):
		return create_sql(
			prov,
			root,
			mem.New,
		)
	case utils.AnyEq([]string{
		"postgres", "pg",
	}, st):
		return create_sql(
			prov,
			root,
			pgres.New,
		)
	case utils.AnyEq([]string{
		"mysql", "maria",
	}, st):
		return create_sql(
			prov,
			root,
			mysql.New,
		)
	}
	return nil, fmt.Errorf("invalid storage type: %s", st)
}

func with_create(config config.ConfigProvider) {
	config.GetConfig().Storage.Create = true
}

func with_logger() {

}
