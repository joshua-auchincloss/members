package service

import (
	"context"
	"members/common"
	"members/config"
	errs "members/errors"
	"members/storage/base"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"go.uber.org/fx"
)

type (
	SvcFramework struct {
		mu     *sync.Mutex
		in     map[common.Service]wrapped
		health func(string) Service
	}
	wrapped = func(health, rpc string) Service
)

var (
	Module = fx.Options(
		fx.Provide(New),
	)
)

func New(prov config.ConfigProvider) *SvcFramework {
	return &SvcFramework{
		mu: new(sync.Mutex),
		in: make(map[common.Service]wrapped),
	}
}

func (s *SvcFramework) Get(key common.Service) wrapped {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.in[key]
}

func (s *SvcFramework) Set(key common.Service, ar wrapped) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.in[key] = ar
}

func (s *SvcFramework) Start(
	lc fx.Lifecycle,
	prov config.ConfigProvider) {
	cf := prov.GetConfig()
	for k, svc := range s.in {
		do := ForService(prov, k)
		log.Info().Str("svc", common.ServiceKeys.Get(k)).Bool("do", do).Msg("HERE")
		if do {
			ports := cf.Members.GetService(k)
			for i, sv := range ports.Svc.Service {
				sp := prov.HostPort(sv)
				hp := prov.HostPort(ports.Svc.Health[i])
				service := svc(hp, sp)
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						log.Printf("starting service %s [shard %d]", common.ServiceKeys.Get(k), i+1)
						log.Printf("service %s (pid: %d)", common.ServiceKeys.Get(k), os.Getpid())
						service.Chain(ctx)
						return service.Start(ctx)
					},
					OnStop: func(ctx context.Context) error {
						return service.Stop(ctx)
					},
				})
			}
		}
	}
}

func Create[T Service](key common.Service) func(
	prov config.ConfigProvider,
	fw *SvcFramework,
	factory ServiceFactory[T],
	store base.BaseStore,
	root *zerolog.Logger,
	watcher errs.Watcher) error {
	log.Info().Str("svc", common.ServiceKeys.Get(key)).Msg("with create")
	return func(
		prov config.ConfigProvider,
		fw *SvcFramework,
		factory ServiceFactory[T],
		store base.BaseStore,
		root *zerolog.Logger,
		watcher errs.Watcher) error {
		cfg := prov.GetConfig()
		ctx := context.Background()
		ctx = context.WithValue(ctx, common.ContextKeyWithStore, store)
		if key == common.ServiceHealth {
			fw.mu.Lock()
			defer fw.mu.Unlock()
			// capture the health factory role
			fw.health = func(health string) Service {
				h := factory.CreateService(ctx, prov)
				h.WithBase(NewBase(
					key,
					prov,
					watcher,
					cfg.Members.Dns,
					health,
					time.Second*10,
				))
				return h
			}
			return nil
		}

		// capture the registry services & pair to a health service
		fu := func(health, rpc string) Service {
			svc := factory.CreateService(ctx, prov)
			svc.WithBase(NewBase(
				key,
				prov,
				watcher,
				cfg.Members.Dns,
				rpc,
				poll_freq,
			))

			// link health chain to the service chain
			h := fw.health(health)
			h.WithChained(svc)

			// link both services
			h.WithLink(svc)
			svc.WithLink(h)

			// build service-aware loggers
			h.BuildLogger(root)
			svc.BuildLogger(root)
			return h
		}

		fw.Set(key, fu)
		return nil
	}
}
