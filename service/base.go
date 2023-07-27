package service

import (
	"context"
	"errors"
	"members/common"
	"members/config"
	errs "members/errors"
	"members/logging"
	"members/utils"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

func noop(context.Context) error { return nil }

type (
	Chain = func(ctx context.Context) error

	BaseChain struct {
		// chain startup
		start Chain
		// chain loop events
		next Chain
		// chain cleanup events
		cleanup Chain
	}

	BaseService struct {
		*BaseChain
		mu *sync.Mutex

		role   common.Service
		status common.Status

		prov   config.ConfigProvider
		logger *zerolog.Logger

		// capture centralized errors
		watcher errs.Watcher
		// local channel to receive errors
		errs chan error
		// local channel to receive signals
		sigs chan os.Signal
		// how often to poll all the above,
		// calling the .next(ctx) [error]
		// chain
		tick time.Duration

		service string
		dns     string

		link Service
	}
)

var (
	poll_freq time.Duration = time.Millisecond * 500
	errDone                 = errors.New("done")
)

func NewBase(
	role common.Service,
	prov config.ConfigProvider,
	watcher errs.Watcher,
	dns, addr string,
	tick time.Duration,
) *BaseService {
	if utils.ZeroStr(dns) {
		dns = addr
	}
	return &BaseService{
		&BaseChain{nil, nil, nil},
		new(sync.Mutex),
		role,
		common.NoStatus,
		prov,
		nil,
		watcher,
		watcher.Subscription(),
		utils.Sigs(),
		tick,
		addr,
		dns,
		nil,
	}
}

func (s *BaseService) Name() string {
	return common.ServiceKeys.Get(s.role)
}

func (s *BaseService) Role() common.Service {
	return s.role
}

func (s *BaseService) Compliment() Service {
	return s.link
}

func (s *BaseService) DNS() string {
	switch s.Role() {
	case common.ServiceHealth:
		return s.Compliment().DNS()
	default:
		return s.dns
	}
}

func (s *BaseService) Address() string {
	return s.service
}

func (s *BaseService) Health() string {
	switch s.role {
	case common.ServiceHealth:
		return s.Address()
	default:
		return s.Compliment().Address()
	}
}

func (s *BaseService) Service() string {
	switch s.role {
	case common.ServiceHealth:
		return s.Compliment().Address()
	default:
		return s.Address()
	}
}

func (s *BaseService) GetErrs() chan error {
	return s.errs
}

func (s *BaseService) GetProv() config.ConfigProvider {
	return s.prov
}

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

func (s *BaseService) LoopedStarter(ctx context.Context,
	clean Chain, chain ...Chain) {
	s.WithNext(chain...)
	s.WithStop(clean)
	s.SetStatus(common.StatusStarted)
	tick := time.NewTicker(s.tick)
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
		<-tick.C
	}
}

func (s *BaseService) GetLogger() *zerolog.Logger {
	return s.logger
}

func (s *BaseService) WithOp(
	root Chain,
	get func(*BaseService) Chain,
	set func(Chain),
	loop ...Chain,
) {
	var chain Chain
	prev := get(s)
	if prev == nil {
		chain = root
	} else {
		chain = prev
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
	set(chain)
}

func (s *BaseService) WithChained(
	svc ...Service,
) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.WithStart(func(ctx context.Context) error {
		for _, sv := range svc {
			if err := sv.Start(ctx); err != nil {
				return err
			}
		}
		return nil
	})
	s.WithStop(
		func(ctx context.Context) error {
			for _, sv := range svc {
				if err := sv.Stop(ctx); err != nil {
					s.errs <- err
				}
			}
			return nil
		},
	)
}

func (s *BaseService) Stop(ctx context.Context) error {
	s.SetStatus(common.StatusClosed)
	return s.cleanup(ctx)
}

func (s *BaseService) BuildLogger(
	root *zerolog.Logger,
) {
	s.mu.Lock()
	defer s.mu.Unlock()
	sub := logging.WithSub(root, s.Name(), func(ctx zerolog.Context) zerolog.Context {
		cfg := s.prov.GetConfig()
		ishealth := s.role == common.ServiceHealth
		if ishealth {
			ctx = ctx.
				Str("address", s.service).
				Str("service-address", s.Compliment().Address())
		} else {
			ctx = ctx.Str("address", s.service).
				Str("health-address", s.Compliment().Address())
		}
		return ctx.
			Str("protocol",
				cfg.Members.Protocol,
			).Bool("health",
			ishealth,
		)
	})
	s.logger = sub
}

func (s *BaseService) Logger() *zerolog.Logger {
	return s.logger
}

func (s *BaseService) loop(ctx context.Context) error {
	if s.status != common.StatusStarted {
		return errDone
	}
	select {
	case <-s.sigs:
		s.logger.Info().Msg("closing")
		return errDone
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
func (s *BaseService) WithNext(loop ...Chain) {
	s.WithOp(
		s.loop,
		func(bs *BaseService) Chain { return s.next },
		func(c Chain) { s.next = c },
		loop...,
	)
}

func (s *BaseService) WithStart(loop ...Chain) {
	s.WithOp(
		noop,
		func(bs *BaseService) Chain { return s.start },
		func(c Chain) { s.start = c },
		loop...,
	)
}

func (s *BaseService) WithStop(loop ...Chain) {
	s.WithOp(
		noop,
		func(bs *BaseService) Chain { return s.cleanup },
		func(c Chain) { s.cleanup = c },
		loop...,
	)
}

func (s *BaseService) Chain(ctx context.Context) error {
	return s.start(ctx)
}

func (s *BaseService) WithLink(o Service) error {
	if s.link != nil {
		return errs.ErrOccupied
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.link = o
	return nil
}
