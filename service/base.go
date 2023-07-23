package service

import (
	"context"
	"members/common"
	"members/config"
	errs "members/errors"
	"sync"
)

type (
	chain       = func(ctx context.Context) error
	BaseService struct {
		prov config.ConfigProvider
		mu   *sync.Mutex
		key  common.Service
		errs chan error

		health  string
		service string

		chain   chain
		cleanup func() error
	}
)

func (s *BaseService) GetKey() common.Service {
	return s.key
}
func (s *BaseService) GetMu() *sync.Mutex {
	return s.mu
}
func (s *BaseService) GetHealth() string {
	return s.health
}
func (s *BaseService) GetService() string {
	return s.service
}
func (s *BaseService) GetErrs() chan error {
	return s.errs
}
func (s *BaseService) GetProv() config.ConfigProvider {
	return s.prov
}

func (s *BaseService) GetBase() *BaseService {
	return s
}

func (s *BaseService) NewBase(
	prov config.ConfigProvider,
	watcher errs.Watcher,
	health, service string,
) *BaseService {
	return &BaseService{prov, new(sync.Mutex), s.key, watcher.Subscription(), health, service, nil, nil}
}

func (s *BaseService) WithChainer(
	svc ...Service,
) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.chain = func(ctx context.Context) error {
		for _, sv := range svc {
			go sv.Start(ctx)
		}
		return nil
	}
	s.cleanup = func() error {
		for _, sv := range svc {
			if err := sv.Stop(); err != nil {
				s.errs <- err
				return err
			}
		}
		return nil
	}
}

func (s *BaseService) Close(ctx context.Context) error {
	return s.cleanup()
}

func (s *BaseService) Chain(ctx context.Context) error {
	return s.chain(ctx)
}

func (s *BaseService) WithKey(key common.Service) {
	s.key = key
}
