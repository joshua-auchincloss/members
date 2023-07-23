package health

import (
	"context"
	"errors"
	"log"
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

		ticker time.Ticker
		status common.Status

		store storage.Store
	}
)

var (
	_ service.Service = ((*healthService)(nil))

	done = errors.New("done")
)

var (
	serving = pkg.HealthCheckResponse{
		Status: pkg.HealthCheckResponse_SERVING,
	}
)

func (h *healthService) setStatus(status common.Status) {
	h.GetMu().Lock()
	h.status = status
	h.GetMu().Unlock()
}

func (h *healthService) WithBase(base service.BaseService) {
	h.BaseService = base
}

func (h *healthService) loop() error {
	if h.status != common.StatusStarted {
		return done
	}
	ctx, cancel := context.WithTimeout(context.TODO(), default_polling/2)
	defer cancel()
	if err := h.store.UpsertMembership(context.TODO(), &common.Membership{
		Service:        h.GetKey(),
		PublicAddress:  h.GetService(),
		JoinTime:       time.Now(),
		LastHealthTime: time.Now(),
	}); err != nil {
		log.Printf("err: %s", err)
		return err
	} else {
		log.Print("upserted")
	}

	if memb, err := h.store.GetMembers(ctx); err != nil {
		log.Print(err)
		return err
	} else {
		for _, mem := range memb {
			log.Printf(
				"%+v", *mem,
			)
		}
	}

	log.Printf("health %d: serving", h.GetKey())
	select {
	case <-ctx.Done():
		return nil
	case <-h.ticker.C:
	}
	return nil
}

func (h *healthService) Start(ctx context.Context) {
	pth, handle := healthconnect.NewHealthHandler(h)
	go service.GrpcStarter(h.GetHealth(), pth, handle)
	if h.status < common.StatusStarted {
		h.setStatus(common.StatusStarted)
		for {
			if err := h.loop(); err != nil {
				log.Print(err)
				h.GetErrs() <- err
				return
			}
		}
	}
}
func (h *healthService) Stop() error {
	log.Print("health stopping")
	h.setStatus(common.StatusClosed)
	return nil
}

func (h *healthService) Check(ctx context.Context, req *connect.Request[pkg.HealthCheckRequest]) (*connect.Response[pkg.HealthCheckResponse], error) {
	return connect.NewResponse(&serving), nil
}

func (h *healthService) Watch(ctx context.Context, req *connect.Request[pkg.HealthCheckRequest], stream *connect.ServerStream[pkg.HealthCheckResponse]) error {
	tick := time.NewTicker(time.Millisecond * 250)
	log.Print(req.Msg.Service)
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
