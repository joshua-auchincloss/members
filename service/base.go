package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"members/common"
	"members/config"
	errs "members/errors"
	"sync"
	"time"
)

type (
	chain       = func(ctx context.Context) error
	BaseService struct {
		prov    config.ConfigProvider
		mu      *sync.Mutex
		key     common.Service
		errs    chan error
		status  common.Status
		health  string
		service string

		chain   chain
		cleanup func() error

		tick *time.Ticker

		ishealth bool
	}

	loop_next = func(context.Context) error
)

var (
	poll_freq time.Duration = time.Millisecond * 500
	done                    = errors.New("done")
)

func (s *BaseService) SetStatus(status common.Status) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.status = status
}

func (s *BaseService) GetStatus() common.Status {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.status
}

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

func (s *BaseService) Name() string {
	svc := service_keys[s.key]
	if s.ishealth {
		return fmt.Sprintf("%s (health)", svc)
	}
	return svc
}

func (s *BaseService) loop(ctx context.Context) error {
	if s.status != common.StatusStarted {
		return done
	}
	select {
	case <-ctx.Done():
		log.Printf("%s closing", s.Name())
		return done
	default:
		return nil
	}
}

func (s *BaseService) LoopedStarter(ctx context.Context, chain ...loop_next) error {
	s.SetStatus(common.StatusStarted)
	var next loop_next
	if len(chain) > 0 {
		next = s.loop
		for n, ch := range chain {
			prev := next
			next = func(ctx context.Context) error {
				log.Printf("level %d", n)
				if err := prev(ctx); err != nil {
					return err
				}
				return ch(ctx)
			}
		}
	} else {
		next = s.loop
	}
	for {
		if err := next(ctx); err != nil {
			log.Fatal(err)
			s.GetErrs() <- err
			return err
		}
		log.Printf("%s running", s.Name())
		<-s.tick.C
	}
}

func (s *BaseService) NewBase(
	prov config.ConfigProvider,
	watcher errs.Watcher,
	health, service string,
	tick time.Duration,
	ishealth bool,
) *BaseService {
	return &BaseService{prov,
		new(sync.Mutex),
		s.key,
		watcher.Subscription(),
		common.NoStatus,
		health,
		service,
		nil, nil,
		time.NewTicker(tick),
		ishealth,
	}
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
	s.SetStatus(common.StatusClosed)
	return s.cleanup()
}

func (s *BaseService) Chain(ctx context.Context) error {
	return s.chain(ctx)
}

func (s *BaseService) WithKey(key common.Service) {
	s.key = key
}
