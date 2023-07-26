package base

import (
	"members/common"
	"members/config"

	"github.com/rs/zerolog"
	"github.com/uptrace/bun"
)

type (
	Convert = func(*bun.DB) BaseStore
)

func NoOpBase[DB any,
	I func(db DB, log *zerolog.Logger, typ common.StorageType) (BaseStore, error),
](typ common.StorageType, ini I) func(DB) (BaseStore, error) {
	return func(d DB) (BaseStore, error) {
		if db, err := ini(d, nil, typ); err != nil {
			return nil, err
		} else {
			return db, nil
		}
	}
}

func PreStart(
	prov config.ConfigProvider,
	store BaseStore,
	root *zerolog.Logger,
) {
	sub := LoggerFor(prov, root)
	store.WithLogger(sub)
	store.WithProvider(prov)
}
