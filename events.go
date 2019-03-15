package assets

import "github.com/ecletus/plug"

type dis struct {
	dis plug.EventDispatcherInterface
}

func (dis dis) OnSyncConfig(cb func(e *PreRepositorySyncEvent)) {
	dis.dis.On(E_PRE_REPOSITORY_SYNC, func(e plug.EventInterface) {
		cb(e.(*PreRepositorySyncEvent))
	})
}

func Dis(d plug.EventDispatcherInterface) *dis {
	return &dis{d}
}
