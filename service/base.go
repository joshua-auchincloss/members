package service

import (
	"context"
	"errors"
	"members/common"
	"members/config"
	errs "members/errors"
	"members/logging"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

type (
	Chain       = func(ctx context.Context) error
	BaseService struct {
		prov    config.ConfigProvider
		mu      *sync.Mutex
		key     common.Service
		watcher errs.Watcher
		errs    chan error

		status  common.Status
		health  string
		service string

		// chain startup
		start Chain
		// chain loop events
		next Chain
		// chain cleanup events
		cleanup Chain

		tick *time.Ticker

		ishealth bool

		logger *zerolog.Logger
	}
)

var (
	poll_freq time.Duration = time.Millisecond * 500
	errDone                 = errors.New("done")
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
	return common.ServiceKeys.Get(s.key)
}

func (s *BaseService) loop(ctx context.Context) error {
	if s.status != common.StatusStarted {
		return errDone
	}
	select {
	case <-ctx.Done():
		s.logger.Info().Msg("closing")
		return errDone
	default:
		return nil
	}
}

/*
Whether to chain the next iteration or start cleaning up the server

At the core, it always checks the status of the server and whether the cleanup signal is received
*/
func (s *BaseService) WithLoop(loop ...Chain) {
	var chain Chain
	if s.next == nil {
		chain = s.loop
	} else {
		chain = s.next
	}
	if len(loop) > 0 {
		for n, next := range loop {
			prev := chain
			chain = func(ctx context.Context) error {
				s.logger.Debug().Int("recursion", n).Send()
				if err := prev(ctx); err != nil {
					return err
				}
				return next(ctx)
			}
		}
	}
	s.next = chain
}

func (s *BaseService) LoopedStarter(ctx context.Context, clean Chain, chain ...Chain) {
	s.SetStatus(common.StatusStarted)
	s.WithLoop(chain...)
	for {
		if err := s.next(ctx); err != nil {
			s.logger.Err(err).Send()
			s.GetErrs() <- err
			if err := s.cleanup(ctx); err != nil {
				s.GetErrs() <- err
			}
			return
		}
		s.logger.Print("running")
		<-s.tick.C
	}
}

func (s *BaseService) Chain(ctx context.Context) error {
	return s.start(ctx)
}

func (s *BaseService) GetLogger() *zerolog.Logger {
	return s.logger
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
		watcher,
		watcher.Subscription(),
		common.NoStatus,
		health,
		service,
		nil, nil, nil,
		time.NewTicker(tick),
		ishealth,
		nil,
	}
}

func (s *BaseService) BuildLogger(
	root *zerolog.Logger,
) {
	s.mu.Lock()
	defer s.mu.Unlock()
	sub := logging.WithSub(root, s.Name(), func(ctx zerolog.Context) zerolog.Context {
		cfg := s.prov.GetConfig()
		if s.ishealth {
			ctx = ctx.Str("address", s.health)
		} else {
			ctx = ctx.Str("address", s.service)
		}
		return ctx.
			Str("protocol",
				cfg.Members.Protocol,
			).Bool("health",
			s.ishealth,
		)
	})
	s.logger = sub
}

func (s *BaseService) WithChainer(
	svc ...Service,
) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.start = func(ctx context.Context) error {
		for _, sv := range svc {
			go sv.Start(ctx)
		}
		return nil
	}
	s.cleanup = func(ctx context.Context) error {
		for _, sv := range svc {
			if err := sv.Stop(ctx); err != nil {
				s.errs <- err
				return err
			}
		}
		return nil
	}
}

func (s *BaseService) Close(ctx context.Context) error {
	defer s.SetStatus(common.StatusClosed)
	return s.cleanup(ctx)
}

func (s *BaseService) WithKey(key common.Service) {
	s.key = key
}
