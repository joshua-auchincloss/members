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
		DNS() string
		Address() string
		Role() common.Service
		Service() string
		Health() string
	}

	Chaining interface {
		// call the chain loop startup process
		// equivalent to .Start(ctx) [error] however we
		// always want to call this
		Chain(ctx context.Context) error

		// with an operation.
		// if get(svc) -> chain == nil:
		//  -> use root chain

		// then, merge each element in the loop chain
		// finally, set the final chain using a fn
		WithOp(
			root Chain,
			get func(*BaseService) Chain,
			set func(Chain),
			loop ...Chain,
		)

		// whether to proceed serving
		WithNext(loop ...Chain)
		// chains the startup events
		WithStart(loop ...Chain)

		// chains the shutdown between services
		WithStop(loop ...Chain)

		// chains the startup between different services
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
