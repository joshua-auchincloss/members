package config

import (
	"members/common"
	"sync"

	"github.com/rs/zerolog/log"
)

type (
	event_chan    = chan *common.Event
	DynamicConfig struct {
		mu *sync.Mutex
		ch event_chan

		subs     []event_chan
		clusters map[common.Service]*DnsList

		prov ConfigProvider
	}
)

func (d *DynamicConfig) start() {
	for c := range d.ch {
		log.Info().Str("kind", common.EventKeys.Get(c.Kind)).Msg("changed")
		for i, sub := range d.subs {
			sub <- c
			log.Debug().
				Str("kind", common.EventKeys.Get(c.Kind)).
				Int("i", i).
				Msg("propogated")
		}
	}
}

func (d *DynamicConfig) automark_available() {
	d.Lock()
	defer d.Unlock(&common.Event{
		Kind: common.EventBulkMembershipAvailability,
	})
	wg := new(sync.WaitGroup)
	conns := d.prov.GetConfig().Members.ConnectionsPerService
	for _, list := range d.clusters {
		wg.Add(1)
		go list.AvailabilityWalk(wg, int(conns))
	}
	wg.Wait()
}

func (d *DynamicConfig) watch_for_availability() {
	ev := d.Subscription()

	for e := range ev {
		switch e.Kind {
		case common.EventMembershipUpdate:
			d.automark_available()
		case common.EventBulkMembershipUpdate:
			d.automark_available()
		}
	}
}

func (d *DynamicConfig) Subscription() event_chan {
	d.mu.Lock()
	defer d.mu.Unlock()
	ch := make(event_chan)
	d.subs = append(d.subs, ch)
	return ch
}

func (d *DynamicConfig) RLock() {
	d.mu.Lock()
}

func (d *DynamicConfig) Lock() {
	d.mu.Lock()
}

func (d *DynamicConfig) Unlock(event *common.Event) {
	d.mu.Unlock()
	d.ch <- event
}

func (d *DynamicConfig) RUnlock() {
	d.mu.Unlock()
}

func (d *DynamicConfig) GetDns(svc common.Service, dns string) []string {
	d.RLock()
	defer d.RUnlock()
	known := []string{}
	aware, ok := d.clusters[svc]
	if !ok {
		return known
	}
	for _, addrs := range aware.candidates(dns) {
		known = append(known, addrs...)
	}
	return known
}
func (d *DynamicConfig) PeekAll() map[common.Service]*DnsList {
	d.RLock()
	defer d.RUnlock()
	return d.clusters
}

func (d *DynamicConfig) AllKnown() map[string][]string {
	d.RLock()
	defer d.RUnlock()
	known := map[string][]string{}
	for _, dns_clusters := range d.clusters {
		for dns, addrs := range dns_clusters.candidates() {
			var r []string
			var ok bool
			if r, ok = known[dns]; !ok {
				r = []string{}
			} else {
				r = append(r, addrs...)
			}
			known[dns] = r
		}
	}
	return known
}
func (d *DynamicConfig) Sync(members []*common.Membership) {
	d.Lock()
	defer d.Unlock(&common.Event{
		Kind: common.EventBulkMembershipUpdate,
	})
	for _, memb := range members {
		// mu will sync this
		d.AddCluster(memb.Service, memb.Dns, memb.PublicAddress, false)
	}
}

func (d *DynamicConfig) AddCluster(svc common.Service, dns, addr string, lock ...bool) {
	l := true
	if len(lock) != 0 {
		l = lock[0]
	}
	if l {
		d.Lock()
		defer d.Unlock(&common.Event{
			Kind: common.EventMembershipUpdate,
		})
	}
	m, ok := d.clusters[svc]
	if !ok {
		m = new_dns(dns)
	}
	m.add_to_cluster(svc, dns, addr, d.prov.GetConfig().Members.Protocol)
	d.clusters[svc] = m
}

func build_client_args(
	svc common.Service,
	args *ClientArgs) *DnsList {
	dns := new_dns(
		args.Dns,
		member_list(common.ServiceAdmin, args.Addresses)...,
	)
	for d, members := range args.Trusted {
		dns.merge_with(common.ServiceAdmin, d, member_list(common.ServiceAdmin, members.Addresses))
	}
	return dns
}

func add_client_to_start(
	svc common.Service,
	known *Members,
	start *map[common.Service]*DnsList,
) {
	cli := known.GetClient(svc)
	if cli != nil {
		(*start)[svc] = build_client_args(svc, cli)
	}
}

func getDynamic(prov ConfigProvider) (*DynamicConfig, error) {
	known := prov.GetConfig().Members
	start := map[common.Service]*DnsList{}
	add_client_to_start(common.ServiceAdmin, known, &start)
	add_client_to_start(common.ServiceRegistry, known, &start)
	log.Info().Interface("start", known).Send()
	dyn := &DynamicConfig{
		mu:       new(sync.Mutex),
		ch:       make(event_chan),
		clusters: start,
		prov:     prov,
	}
	go dyn.start()
	go dyn.watch_for_availability()
	dyn.automark_available()
	return dyn, nil
}
