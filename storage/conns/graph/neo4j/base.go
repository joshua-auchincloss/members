package neo4j

import (
	"context"
	"members/common"
	"members/config"
	"members/storage/base"
	"time"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/rs/zerolog"
)

type (
	Neo4jBase struct {
		DB       neo4j.DriverWithContext
		logger   *zerolog.Logger
		kind     common.StorageType
		prov     config.ConfigProvider
		sess_cfg *neo4j.SessionConfig
	}
	GraphInitializer = func(prov config.ConfigProvider) (neo4j.DriverWithContext, error)
	GraphConverter   = func(neo4j.DriverWithContext) (base.BaseStore, error)
)

var (
	_ base.BaseStore = ((*Neo4jBase)(nil))
)

func maybe[T any](rec *neo4j.Record, key string) T {
	r, ok := rec.Get(key)
	if !ok {
		return *new(T)
	}
	return r.(T)
}

func NewNeo4jBase(db neo4j.DriverWithContext, log *zerolog.Logger, typ common.StorageType) (base.BaseStore, error) {
	return &Neo4jBase{db, log, typ, nil, nil}, nil
}

func (gr *Neo4jBase) session(ctx context.Context) neo4j.SessionWithContext {
	return gr.DB.NewSession(ctx, *gr.sess_cfg)
}

func (gr *Neo4jBase) build_cfg() {
	gr.sess_cfg = &neo4j.SessionConfig{
		DatabaseName: gr.prov.GetConfig().Storage.DB,
	}
}

func (gr *Neo4jBase) Kind() common.StorageType {
	return gr.kind
}

func (gr *Neo4jBase) WithQueryHook(i interface{}) {
	gr.DB.NewSession(context.TODO(), neo4j.SessionConfig{})
}

func (gr *Neo4jBase) WithLogger(sub *zerolog.Logger) {
	gr.logger = sub
}

func (gr *Neo4jBase) WithProvider(prov config.ConfigProvider) {
	gr.prov = prov
	gr.build_cfg()
}

func (gr *Neo4jBase) Logger() *zerolog.Logger {
	return gr.logger
}

func (gr *Neo4jBase) UpsertMembership(ctx context.Context, meta *common.Membership) error {
	_, err := gr.session(ctx).Run(ctx, `
		merge (
			m:Member { address: $address, service: $service }
		)
		set m.dns = $dns
		set m.registration = $registration
		set m.last_health = $last_health
	`, map[string]any{
		"dns":          meta.Dns,
		"address":      meta.PublicAddress,
		"service":      meta.Service,
		"registration": meta.JoinTime.UTC(),
		"last_health":  meta.LastHealthTime.UTC(),
	})
	return err
}

func (gr *Neo4jBase) CleanOldMembers(ctx context.Context, from time.Duration) error {
	return nil
}

func (gr *Neo4jBase) GetMembers(ctx context.Context, kind ...common.Service) ([]*common.Membership, error) {
	res, err := gr.session(ctx).Run(ctx, `
	match (m:Member)
	where m.last_health + duration('30s') => now()
	return m
	`, map[string]any{})
	if err != nil {
		return nil, err
	}
	rec, err := res.Collect(ctx)
	if err != nil {
		return nil, err
	}
	out := []*common.Membership{}
	for _, rec := range rec {
		out = append(out, &common.Membership{
			Dns:            maybe[string](rec, "dns"),
			PublicAddress:  maybe[string](rec, "address"),
			Service:        maybe[common.Service](rec, "service"),
			JoinTime:       maybe[time.Time](rec, "registration"),
			LastHealthTime: maybe[time.Time](rec, "last_health"),
		})
	}
	return out, nil
}

func (gr *Neo4jBase) Teardown(ctx context.Context) error {
	return nil
}

func (gr *Neo4jBase) Setup(ctx context.Context, cfg config.ConfigProvider) error {
	storg := cfg.GetConfig().Storage
	if storg.Drop {
		sess := gr.session(ctx)
		_, err := sess.Run(ctx, `
		match (n) detach delete n
		`, map[string]any{})
		if err != nil {
			return err
		}
		_, err = sess.Run(ctx, `
		drop index $index
		`, map[string]any{
			"index": common.MembershipIndex,
		})
		if err != nil {
			return err
		}
	}
	if storg.Create {
		sess := gr.session(ctx)
		result, err := sess.Run(ctx, `
		create constraint $index
			for (member:Member)
		require (member.address, member.service)
		is unique
		`, map[string]any{
			"index": common.MembershipIndex,
		})
		if err != nil {
			return err
		}
		rec, err := result.Collect(ctx)
		if err != nil {
			return err
		}
		for _, rec := range rec {
			gr.logger.Info().Interface("record", rec).Send()
		}
	}
	return nil
}
