package service

import (
	"context"
	"members/common"
	"members/config"
)

type (
	Lifecycle interface {
		Start(ctx context.Context) error
		Stop(ctx context.Context) error
	}

	ServiceMeta interface {
		GetDns() string
		GetKey() common.Service
		GetHealth() string
		GetService() string
		WithKey(key common.Service)
	}

	Chaining interface {
		Chain(ctx context.Context) error

		WithOp(
			root Chain,
			get func(*BaseService) Chain,
			set func(Chain),
			loop ...Chain,
		)

		WithNext(loop ...Chain)
		WithStart(loop ...Chain)
		WithStop(loop ...Chain)

		WithChained(
			svc ...Service,
		)
	}

	Service interface {
		Lifecycle
		ServiceMeta
		Chaining
		common.Logs

		WithBase(*BaseService)
		Compliment() Service
		WithLink(Service) error
	}

	ServiceFactory[T Service] interface {
		CreateService(ctx context.Context, cfg config.ConfigProvider) T
	}
)
