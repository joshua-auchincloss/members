package errs

import (
	"github.com/rs/zerolog/log"

	"go.uber.org/fx"
)

type (
	watcher struct {
		ch chan error
	}

	Watcher interface {
		Subscription() chan error
	}
)

var (
	Module = fx.Module(
		"error-watcher",
		fx.Provide(new),
	)
)

func new() Watcher {
	w := watcher{make(chan error)}
	go w.start()
	return &w
}

func (w *watcher) start() {
	for err := range w.ch {
		log.Printf("err: %s", err)
	}
}

func (w *watcher) Subscription() chan error {
	ch := make(chan error)
	go func(ch chan error) {
		for err := range ch {
			w.ch <- err
		}
	}(ch)
	return ch
}
