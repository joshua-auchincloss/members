package cluster

import (
	"members/utils"

	"github.com/rs/zerolog/log"
)

func filter_list_for[T any](m DnsList,
	filt func(addr ClusterMember) bool,
	get func(addr ClusterMember) T,
	dns ...string,
) map[string][]T {
	var d string
	if len(dns) != 0 {
		d = dns[0]
	}
	check := !utils.ZeroStr(d)
	known := map[string][]T{}

	for dns, addrs := range m.Peek() {
		var this bool
		if check {
			this = dns == d
		} else {
			this = true
		}
		if this {
			ar, ok := known[dns]
			if !ok {
				ar = []T{}
			}
			for _, addr := range addrs {
				ok := filt(addr)
				if ok {
					ar = append(ar, get(addr))
				}
				log.Trace().
					Str("dns", dns).
					Interface("member", addr).
					Bool("met", ok).
					Msg("filter checked")
			}
			known[dns] = append(ar, ar...)
		}
	}
	log.Debug().Interface("known", known).Msg("filtering done")
	return known
}
