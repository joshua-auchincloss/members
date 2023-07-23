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

		logger *zerolog.Logger
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
	return service_keys[s.key]
}

func (s *BaseService) loop(ctx context.Context) error {
	if s.status != common.StatusStarted {
		return done
	}
	select {
	case <-ctx.Done():
		s.logger.Info().Msg("closing")
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
				s.logger.Debug().Msgf("level %d", n)
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
			s.logger.Err(err).Send()
			s.GetErrs() <- err
			return err
		}
		s.logger.Print("running")
		<-s.tick.C
	}
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
		watcher.Subscription(),
		common.NoStatus,
		health,
		service,
		nil, nil,
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
		ctx = ctx.Bool("health", s.ishealth)
		if s.ishealth {
			ctx = ctx.Str("address", s.health)
		} else {
			ctx = ctx.Str("address", s.service)
		}
		return ctx
	})
	s.logger = sub
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
