package cluster

import (
	"members/common"
	"sync"
)

func MemberList(role common.Service, li []string) []ClusterMember {
	members := make([]ClusterMember, len(li))
	for i, member := range li {
		members[i] = &clusterMember{
			mu:   new(sync.Mutex),
			addr: member,
			role: role,
		}
	}
	return members
}

func Addresses(li []ClusterMember) []string {
	members := make([]string, len(li))
	for i, member := range li {
		members[i] = member.Address()
	}
	return members
}
