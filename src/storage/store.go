package storage

import (
	"context"
	"members/common"
	"members/config"
	wctx "members/context"
)

type (
	Store interface {
		Setup(config config.ConfigProvider) error
		Teardown() error
		Registered(key string) bool
		GetHandler(key string) (*common.RegisteredProto, error)
		UpsertMembership(ctx context.Context, meta *common.Membership) error
		GetMembers(ctx context.Context, kind ...common.Service) ([]*common.Membership, error)
		CreateProject(ctx context.Context, project *common.ProtoProject, proto *common.ProtoMeta) error
		RegisterProto(ctx context.Context, proto *common.ProtoMeta, data *common.RegisteredProto) error
	}

	StorageType = int
)

func Setup(config config.ConfigProvider, store Store) error {
	return store.Setup(config)
}

func WithStore(store Store) wctx.Context {
	return context.WithValue(wctx.New(), wctx.ContextKeyWithStore, store)
}

func GetStore(ctx wctx.Context) Store {
	return ctx.Value(wctx.ContextKeyWithStore).(Store)
}

const (
	Memory StorageType = iota
	Postgres
	Mysql
)
