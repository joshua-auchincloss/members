package health

import (
	"members/common"
	"members/config"
	"members/service"
	"members/storage"
	"time"

	"go.uber.org/fx"
)

type (
	healthFactory struct {
		poll time.Duration
	}
	ServiceFactory = service.ServiceFactory[*healthService]
)

var (
	Module = fx.Module(
		"health-service",
		fx.Provide(
			fx.Annotate(
				New,
				fx.As(new(ServiceFactory)),
			),
		),
		fx.Invoke(
			service.Create[*healthService](common.ServiceHealth),
		),
	)
)

var (
	_ ServiceFactory = ((*healthFactory)(nil))

	default_polling = time.Second * 5
)

func New() ServiceFactory {
	return &healthFactory{
		default_polling,
	}
}

func (h *healthFactory) CreateService(cfg config.ConfigProvider, store storage.Store) *healthService {
	return &healthService{
		ticker: *time.NewTicker(h.poll),
		status: common.NoStatus,
		store:  store,
	}
}
