package assets

import (
	"github.com/aghape/aghape"
	"github.com/aghape/plug"
	"github.com/moisespsena-go/sortvalues"
	"github.com/moisespsena/go-assetfs"
	"github.com/moisespsena/go-assetfs/repository/api"
	"github.com/moisespsena/go-path-helpers"
)

var (
	pkg                   = path_helpers.GetCalledDir()
	E_PRE_REPOSITORY_SYNC = pkg + ".repo_sync"
	CB_ASSETS_SYNC_CONFIG = pkg + ".repo_sync_config"
)

type PreRepositorySyncEvent struct {
	plug.PluginEventInterface
	Repo api.Interface
}

func SyncConfigCallbackFactory(repo api.Interface, agp *aghape.Aghape) func() {
	return func() {
		e := &PreRepositorySyncEvent{plug.NewPluginEvent(E_PRE_REPOSITORY_SYNC), repo}
		if err := agp.Plugins().TriggerPlugins(e); err != nil {
			panic(err)
		}
	}
}

func SyncConfigCallbackRegister(register func(f func(), name string) (v sortvalues.ValueInterface), repo api.Interface, agp *aghape.Aghape) {
	register(SyncConfigCallbackFactory(repo, agp), CB_ASSETS_SYNC_CONFIG).
		After(assetfs.CB_REPO_CONFIG_INIT).
		Before(assetfs.CB_REPO_SYNC)
}