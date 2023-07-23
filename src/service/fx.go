package service

import (
	"context"
	"log"
	"members/common"
	"members/config"
	errs "members/errors"
	"members/storage"
	"os"
	"sync"

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
		health func() Service
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
		if ForService(prov, k) {
			ports := cf.Members.GetService(k)
			for i, sv := range ports.Service {
				sp := prov.HostPort(sv)
				hp := prov.HostPort(ports.Health[i])
				service := svc(hp, sp)
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						log.Printf("starting service %s [shard %d]", service_keys[k], i+1)
						log.Printf("service %s (pid: %d)", service_keys[k], os.Getpid())
						service.Chain(ctx)
						go service.Start(ctx)
						return nil
					},
					OnStop: func(ctx context.Context) error {
						err := service.GetBase().Close(ctx)
						if err != nil {
							log.Printf("error closing children: %s", err)
						}
						return service.Stop()
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
	watcher errs.Watcher) error {
	return func(
		prov config.ConfigProvider,
		fw *SvcFramework,
		factory ServiceFactory[T],
		store storage.Store,
		watcher errs.Watcher) error {
		if key == common.ServiceHealth {
			fw.mu.Lock()
			fw.health = func() Service {
				return factory.CreateService(prov, store)
			}
			defer fw.mu.Unlock()
			return nil
		}
		fu := func(health, rpc string) Service {
			svc := factory.CreateService(prov, store)
			svc.WithKey(key)
			svc.WithBase(*svc.NewBase(prov, watcher, health, rpc))
			h := fw.health()
			h.WithKey(key)
			h.WithBase(*h.NewBase(prov, watcher, health, rpc))
			h.WithChainer(svc)
			return h
		}

		fw.Set(key, fu)
		return nil
	}
}
