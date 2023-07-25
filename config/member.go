package config

import (
	"members/common"
	"members/utils"
	"net"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

type (
	ClusterMember struct {
		mu *sync.Mutex

		Role  common.Service
		Proto string
		Addr  string

		checked bool
		visible bool
	}

	DnsList struct {
		mu      *sync.Mutex
		members map[string][]*ClusterMember
	}

	fl[T any] struct {
		k string
		m []T
	}
)

func WalkDnsList(l *DnsList,
	do func(addr *ClusterMember) error,
	cap int,
	dns ...string) error {
	var d string
	if len(dns) != 0 {
		d = dns[0]
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	check := !utils.ZeroStr(d)
	max := cap != -1 && cap > 1
	var seen int
	for dns, addrs := range l.members {
		var this bool
		if check {
			this = dns == d
		} else {
			this = true
		}
		if this {
			for _, addr := range addrs {
				if max {
					seen += 1
					if seen == cap {
						log.Warn().Int("cap", cap).Msg("cap reached")
						return nil
					}
					log.Debug().Int("seen", seen).Int("cap", cap).Send()
				}
				if err := do(addr); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (m *ClusterMember) Available() bool {
	m.mu.Lock()
	return m.visible
}

func (d *DnsList) Peek() map[string][]*ClusterMember {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.members
}

func (d *DnsList) AvailabilityWalk(wg *sync.WaitGroup, depth int) {
	defer wg.Done()
	if err := WalkDnsList(
		d,
		func(addr *ClusterMember) error {
			if !addr.checked {
				addr.mu.Lock()
				defer addr.mu.Unlock()
				log.Info().Str("proto", "tcp").Str("address", addr.Addr).Send()
				if conn, err := net.DialTimeout("tcp", addr.Addr, time.Millisecond*1); err != nil {
					switch err {
					case net.ErrClosed:
					default:
					}
					addr.visible = false
				} else {
					addr.visible = true
					defer conn.Close()
				}
				addr.checked = true
			}
			return nil
		},
		depth,
	); err != nil {
		log.Error().Err(err).Send()
	}
}

func new_dns(dns string, members ...*ClusterMember) *DnsList {
	start := map[string][]*ClusterMember{}
	if len(members) != 0 {
		start[dns] = members
	}
	return &DnsList{
		new(sync.Mutex),
		start,
	}
}

func member_list(role common.Service, li []string) []*ClusterMember {
	members := make([]*ClusterMember, len(li))
	for i, member := range li {
		members[i] = &ClusterMember{
			mu:   new(sync.Mutex),
			Addr: member,
			Role: role,
		}
	}
	return members
}

func addresses(li []*ClusterMember) []string {
	members := make([]string, len(li))
	for i, member := range li {
		members[i] = member.Addr
	}
	return members
}

func filter_list_for[T any](m *DnsList,
	filt func(addr *ClusterMember) bool,
	get func(addr *ClusterMember) T,
	dns ...string,
) map[string][]T {
	var d string
	if len(dns) != 0 {
		d = dns[0]
	}
	check := !utils.ZeroStr(d)
	known := map[string][]T{}

	for dns, addrs := range m.members {
		var this bool
		if check {
			this = dns == d
		} else {
			this = true
		}
		if this {
			log.Info().Str("dns", dns).Interface("member", addrs).Msg("checking")
			ar, ok := known[dns]
			if !ok {
				ar = []T{}
			}
			for _, addr := range addrs {
				if filt(addr) {
					ar = append(ar, get(addr))
				}
			}
			f := fl[T]{
				dns,
				ar,
			}
			known[f.k] = append(ar, f.m...)
		}
	}
	log.Debug().Interface("known", known).Msg("filtering done")
	return known
}

func (m *DnsList) candidates(dns ...string) map[string][]string {
	return filter_list_for(
		m,
		func(addr *ClusterMember) bool { return (!addr.checked || addr.visible) },
		func(addr *ClusterMember) string { return addr.Addr },
		dns...,
	)
}

func (m *DnsList) merge_with(svc common.Service, dns string, other []*ClusterMember) {
	m.mu.Lock()
	defer m.mu.Unlock()
	var resolved []*ClusterMember
	var ok bool
	if resolved, ok = m.members[dns]; !ok {
		resolved = other
	} else {
		resolved = append(resolved, other...)
	}
	m.members[dns] = resolved
}
func (m *DnsList) add_to_cluster(svc common.Service, dns, proto, addr string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	c, ok := m.members[dns]
	if !ok {
		c = []*ClusterMember{}
	}
	if !utils.OneOf(c, func(a *ClusterMember) bool {
		return a.Addr == addr
	}) {
		c = append(c, &ClusterMember{
			new(sync.Mutex),
			svc,
			proto,
			addr,
			false,
			false,
		})
	}
	m.members[dns] = c
}
