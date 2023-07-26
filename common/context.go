package common

import (
	"context"
	"members/common/internal"
	"members/utils"
	"time"
)

type (
	ContextKeys = int
)

const (
	ContextKeyWithStore = iota + internal.ContextOffset
	ContextKeyDatabaseName
)

func Deadline(
	t ...time.Duration,
) (context.Context, func()) {
	return context.WithTimeout(context.TODO(), utils.Spread(t, time.Second))
}
