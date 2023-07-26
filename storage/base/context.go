package base

import (
	"context"
	"members/common"
	"members/config"
)

func DefaultContext(prov config.ConfigProvider) func() context.Context {
	return func() context.Context {
		return context.WithValue(context.TODO(), common.ContextKeyDatabaseName, prov.GetConfig().Storage.DB)
	}
}
func WithParent(prov config.ConfigProvider) func(context.Context) context.Context {
	return func(parent context.Context) context.Context {
		return context.WithValue(parent, common.ContextKeyDatabaseName, prov.GetConfig().Storage.DB)
	}
}
