package health

import (
	"context"
	"members/common"
	"members/grpc/api/v1/health/healthconnect"
	health "members/grpc/api/v1/health/healthconnect"
	"members/grpc/api/v1/health/pkg"
	"members/service"
	"members/storage"
	"time"

	"github.com/bufbuild/connect-go"
)

type (
	healthService struct {
		health.UnimplementedHealthHandler
		service.BaseService

		store storage.Store
	}
)

var (
	_ service.Service = ((*healthService)(nil))
)

var (
	serving = pkg.HealthCheckResponse{
		Status: pkg.HealthCheckResponse_SERVING,
	}
)

func (h *healthService) WithBase(base service.BaseService) {
	h.BaseService = base
}

func (h *healthService) loop(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, default_polling/2)
	defer cancel()
	memb := &common.Membership{
		Service:        h.GetKey(),
		PublicAddress:  h.GetService(),
		JoinTime:       time.Now(),
		LastHealthTime: time.Now(),
	}

	if err := h.store.UpsertMembership(ctx, memb); err != nil {
		h.GetLogger().Error().Err(err).Msg("could not upsert membership")
		return err
	} else {
		h.GetLogger().Info().Msgf("upserted: %+v", memb)
	}
	if memb, err := h.store.GetMembers(ctx); err != nil {
		h.GetLogger().Print(err)
		return err
	} else {
		for _, mem := range memb {
			h.GetLogger().Printf(
				"%+v", *mem,
			)
		}
	}
	return nil
}

func (h *healthService) Start(ctx context.Context) error {
	pth, handle := healthconnect.NewHealthHandler(h)
	clean, err := h.GrpcStarter(h.GetHealth(), pth, handle)
	if err != nil {
		return err
	}
	go h.LoopedStarter(
		ctx,
		clean,
		func(ctx context.Context) error {
			return h.loop(ctx)
		},
	)
	return nil
}
func (h *healthService) Stop(ctx context.Context) error {
	h.GetLogger().Print("health stopping")
	return nil
}

func (h *healthService) Check(ctx context.Context, req *connect.Request[pkg.HealthCheckRequest]) (*connect.Response[pkg.HealthCheckResponse], error) {
	return connect.NewResponse(&serving), nil
}

func (h *healthService) Watch(ctx context.Context, req *connect.Request[pkg.HealthCheckRequest], stream *connect.ServerStream[pkg.HealthCheckResponse]) error {
	tick := time.NewTicker(time.Millisecond * 250)
	h.GetLogger().Print(req.Msg.Service)
	for {
		if err := stream.Send(
			&serving,
		); err != nil {
			return err
		}
		select {
		case <-ctx.Done():
			return nil
		case <-tick.C:
		}
	}
}
