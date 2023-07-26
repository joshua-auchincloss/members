package storage

import (
	"context"
	"members/common"
	"members/config"
	"members/storage/base"
	"members/storage/conns/sql"

	"github.com/rs/zerolog"
	"github.com/uptrace/bun"
)

type (
	Store interface {
		base.BaseStore
		Registered(key string) bool
		GetHandler(key string) (*common.RegisteredProto, error)
		CreateProject(ctx context.Context, project *common.ProtoProject, proto *common.ProtoMeta) error
		RegisterProto(ctx context.Context, proto *common.ProtoMeta, data *common.RegisteredProto) error
	}
	storeFactory[
		DB any,
		T func(prov config.ConfigProvider) (DB, error),
		F func(DB) (base.BaseStore, error),
	] struct {
		prov    config.ConfigProvider
		root    *zerolog.Logger
		new     T
		convert F
	}
)

var (
	_ base.StoreFactory = ((*storeFactory[
		*bun.DB,
		sql.SqlInitializer,
		sql.SqlConverter,
	])(nil))
)

func NewStoreFactory[
	DB any,
	T func(prov config.ConfigProvider) (DB, error),
	F func(DB) (base.BaseStore, error),
](
	prov config.ConfigProvider,
	root *zerolog.Logger,
	new T,
	convert F,
) base.StoreFactory {
	return &storeFactory[DB, T, F]{
		prov,
		root,
		new,
		convert,
	}
}

func (s *storeFactory[DB, y, z]) New() (base.BaseStore, error) {
	db, err := s.new(s.prov)
	if err != nil {
		return nil, err
	}
	if impl, err := s.convert(db); err != nil {
		return nil, err
	} else {
		return impl, nil
	}
}
