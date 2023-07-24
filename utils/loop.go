package utils

import (
	"context"
	"errors"
	"time"
)

func LoopOrCancel[T any](
	ctx context.Context,
	max_time time.Duration,
	poll time.Duration,
	do func() (T, error)) (*T, error) {
	tick := time.NewTicker(poll)
	ctx, cancel := context.WithTimeout(ctx, max_time)
	defer cancel()
	res := make(chan *T, 1)
	errs := make(chan error, 1)
	go func() {
		t, err := do()
		res <- &t
		errs <- err
	}()
	for {
		select {
		case t := <-res:
			if t == nil {
				err := <-errs
				return nil, err
			}
			return t, nil
		case err := <-errs:
			return nil, err
		case <-ctx.Done():
			return nil, errors.New("timed out")
		case <-tick.C:
		}
	}
}
