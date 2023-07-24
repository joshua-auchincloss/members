package config

import (
	"members/common"
	"sync"

	"github.com/rs/zerolog/log"
)

type (
	DynamicConfig struct {
		mu *sync.Mutex
		ch chan bool

		subs     []chan bool
		clusters map[common.Service]*DnsList
	}
)

func (d *DynamicConfig) start() {
	for c := range d.ch {
		for _, sub := range d.subs {
			sub <- c
		}
		log.Info().Msg("changed")
	}
}

func (d *DynamicConfig) Subscription() chan bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	ch := make(chan bool)
	d.subs = append(d.subs, ch)
	return ch
}

func (d *DynamicConfig) Lock() {
	d.mu.Lock()
}

func (d *DynamicConfig) Unlock() {
	d.mu.Unlock()
	d.ch <- true
}

func (d *DynamicConfig) GetDns(svc common.Service, dns string) []string {
	d.Lock()
	defer d.Unlock()
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
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.clusters
}

func (d *DynamicConfig) AllKnown() map[string][]string {
	d.Lock()
	defer d.Unlock()
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
	defer d.Unlock()
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
		defer d.Unlock()
	}
	m, ok := d.clusters[svc]
	if !ok {
		m = new_dns(dns)
	}
	m.add_to_cluster(svc, dns, addr)
	d.clusters[svc] = m
}

func build_client_args(
	svc common.Service,
	args *ClientArgs) *DnsList {
	dns := new_dns(
		args.Dns,
		member_list(common.ServiceAdmin, args.Addresses)...,
	)
	for d, members := range args.Servers {
		dns.merge_with(common.ServiceAdmin, d, member_list(common.ServiceAdmin, members))
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
		ch:       make(chan bool),
		clusters: start,
	}
	go dyn.start()
	return dyn, nil
}
