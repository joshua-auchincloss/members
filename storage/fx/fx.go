package storage_fx

import (
	"context"
	"members/config"
	"members/storage/base"

	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

var (
	Dependencies = fx.Module(
		"storage",
		fx.Provide(
			TryDefaults,
			GetStore,
		),
		fx.Invoke(
			base.PreStart,
			SetupStore,
		),
	)

	Setup = fx.Module(
		"setup",
		fx.Invoke(WithCreate),
		Dependencies,
	)
)

func TryDefaults(
	prov config.ConfigProvider,
	root *zerolog.Logger,
) (base.StoreFactory, error) {
	if db, err := New(prov, root, Sql); err != nil {
	} else {
		return db, nil
	}
	if db, err := New(prov, root, DGraph); err != nil {
	} else {
		return db, nil
	}
	if db, err := New(prov, root, Neo4j); err != nil {
		return nil, err
	} else {
		return db, nil
	}
}

func GetStore(prov config.ConfigProvider, factory base.StoreFactory, root *zerolog.Logger) (base.BaseStore, error) {
	return factory.New()
}

func WithCreate(config config.ConfigProvider) {
	config.GetConfig().Storage.Create = true
}

func SetupStore(store base.BaseStore, prov config.ConfigProvider) error {
	return store.Setup(context.TODO(), prov)
}
