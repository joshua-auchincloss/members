package storage

import (
	"context"
	"members/common"
	"members/config"
	"time"

	"github.com/rs/zerolog"
)

type (
	Store interface {
		WithLogger(sub *zerolog.Logger)
		Setup(config config.ConfigProvider) error
		Teardown() error
		Registered(key string) bool
		GetHandler(key string) (*common.RegisteredProto, error)
		UpsertMembership(ctx context.Context, meta *common.Membership) error
		GetMembers(ctx context.Context, kind ...common.Service) ([]*common.Membership, error)
		CleanOldMembers(ctx context.Context, from time.Duration) error
		CreateProject(ctx context.Context, project *common.ProtoProject, proto *common.ProtoMeta) error
		RegisterProto(ctx context.Context, proto *common.ProtoMeta, data *common.RegisteredProto) error
	}

	StorageType = int
)

func Setup(config config.ConfigProvider, store Store) error {
	return store.Setup(config)
}

func WithStore(ctx context.Context, store Store) context.Context {
	return context.WithValue(ctx, common.ContextKeys(common.ContextKeyWithStore), store)
}

func GetStore(ctx context.Context) Store {
	return ctx.Value(common.ContextKeyWithStore).(Store)
}

const (
	Memory StorageType = iota
	Postgres
	Mysql
)
