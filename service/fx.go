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
		health wrapped
	}
	wrapped = func(health, rpc string) Service
)

var (
	Module = fx.Options(
		fx.Provide(frame),
	)
)

func frame(prov config.ConfigProvider) *SvcFramework {
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
	return func(
		prov config.ConfigProvider,
		fw *SvcFramework,
		factory ServiceFactory[T],
		store base.BaseStore,
		root *zerolog.Logger,
		watcher errs.Watcher) error {
		log.Info().Str("svc", common.ServiceKeys.Get(key)).Msg("with create")
		cfg := prov.GetConfig()
		ctx := context.Background()
		ctx = context.WithValue(ctx, common.ContextKeyWithStore, store)
		if key == common.ServiceHealth {
			fw.mu.Lock()
			fw.health = func(health, rpc string) Service {
				h := factory.CreateService(ctx, prov)
				h.WithBase(NewBase(
					common.ServiceHealth,
					prov,
					watcher,
					cfg.Members.Dns,
					health,
					rpc,
					time.Second*10,
					true))
				return h
			}
			defer fw.mu.Unlock()
			return nil
		}
		fu := func(health, rpc string) Service {
			svc := factory.CreateService(ctx, prov)
			svc.WithBase(NewBase(
				key,
				prov,
				watcher,
				cfg.Members.Dns,
				health,
				rpc,
				poll_freq,
				false))
			svc.WithKey(key)
			svc.BuildLogger(root)

			h := fw.health(health, rpc)
			// svc.WithLink(h)
			h.WithKey(key)
			h.BuildLogger(root)
			h.WithChained(svc)
			// h.WithLink(svc)
			return h
		}

		fw.Set(key, fu)
		return nil
	}
}
