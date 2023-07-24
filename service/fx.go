package service

import (
	"context"
	"members/common"
	"members/config"
	errs "members/errors"
	"members/storage"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"go.uber.org/fx"
)

func NewAnnotation() {
}

type (
	AppResult struct {
		Svc     Service
		PID     int
		Service string
		Health  string
	}
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
			for i, sv := range ports.Service {
				sp := prov.HostPort(sv)
				hp := prov.HostPort(ports.Health[i])
				service := svc(hp, sp)
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						log.Printf("starting service %s [shard %d]", common.ServiceKeys.Get(k), i+1)
						log.Printf("service %s (pid: %d)", common.ServiceKeys.Get(k), os.Getpid())
						service.Chain(ctx)
						return service.Start(ctx)
					},
					OnStop: func(ctx context.Context) error {
						err := service.GetBase().Close(ctx)
						if err != nil {
							log.Printf("error closing children: %s", err)
						}
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
	store storage.Store,
	root *zerolog.Logger,
	watcher errs.Watcher) error {
	return func(
		prov config.ConfigProvider,
		fw *SvcFramework,
		factory ServiceFactory[T],
		store storage.Store,
		root *zerolog.Logger,
		watcher errs.Watcher) error {
		log.Info().Str("svc", common.ServiceKeys.Get(key)).Msg("with create")
		cfg := prov.GetConfig()
		if key == common.ServiceHealth {
			fw.mu.Lock()
			fw.health = func(health, rpc string) Service {
				h := factory.CreateService(prov, store)
				h.WithBase(*h.NewBase(prov,
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
			svc := factory.CreateService(prov, store)
			svc.WithBase(*svc.NewBase(prov,
				watcher,
				cfg.Members.Dns,
				health,
				rpc,
				poll_freq,
				false))
			svc.WithKey(key)
			svc.BuildLogger(root)

			h := fw.health(health, rpc)
			h.WithKey(key)
			h.BuildLogger(root)
			h.WithChainer(svc)
			return h
		}

		fw.Set(key, fu)
		return nil
	}
}
