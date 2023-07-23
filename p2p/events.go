package p2p

import (
	"github.com/rs/zerolog/log"

	"github.com/hashicorp/memberlist"
)

type (
	eventDelegate struct {
	}
)

var (
	_ memberlist.EventDelegate = ((*eventDelegate)(nil))
)

func (e *eventDelegate) NotifyJoin(node *memberlist.Node) {
	log.Print(node)
}
func (e *eventDelegate) NotifyLeave(node *memberlist.Node) {
	log.Print(node)
}
func (e *eventDelegate) NotifyUpdate(node *memberlist.Node) {
	log.Print(node)
}
