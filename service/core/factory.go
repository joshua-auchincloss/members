package core

import (
	"context"
	"members/common"
	"members/config"
	"members/storage/base"
)

type (
	serviceFactory[T any] struct {
		create func(cfg config.ConfigProvider, store base.BaseStore) T
	}
)

func New[T any](
	create func(cfg config.ConfigProvider,
		store base.BaseStore) T,
) *serviceFactory[T] {
	return &serviceFactory[T]{
		create,
	}
}

func (h *serviceFactory[T]) CreateService(ctx context.Context, cfg config.ConfigProvider) T {
	store := ctx.Value(common.ContextKeyWithStore).(base.BaseStore)
	return h.create(cfg, store)
}
