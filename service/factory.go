package service

import (
	"context"
	"members/common"
	"members/config"
	errs "members/errors"
	"members/storage"
	"time"
)

type (
	Service interface {
		Start(ctx context.Context)
		Stop() error
		Chain(ctx context.Context) error
		WithChainer(
			svc ...Service,
		)
		GetHealth() string
		GetService() string
		WithKey(key common.Service)
		GetBase() *BaseService
		WithBase(base BaseService)
		NewBase(
			cfg config.ConfigProvider,
			watcher errs.Watcher,
			health, service string,
			tick time.Duration,
			ishealth bool,
		) *BaseService
	}
	ServiceFactory[T Service] interface {
		CreateService(cfg config.ConfigProvider, store storage.Store) T
	}
)
