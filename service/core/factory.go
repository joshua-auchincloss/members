package core

import (
	"members/config"
	"members/storage/base"
)

type (
	serviceFactory[T any] struct {
		create func(cfg config.ConfigProvider, store base.BaseStore) T
	}
)

func New[T any](
	create func(cfg config.ConfigProvider, store base.BaseStore) T,
) *serviceFactory[T] {
	return &serviceFactory[T]{
		create,
	}
}

func (h *serviceFactory[T]) CreateService(cfg config.ConfigProvider, store base.BaseStore) T {
	return h.create(cfg, store)
}
