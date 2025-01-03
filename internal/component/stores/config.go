/*
####### dolmen (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package stores

import (
	"github.com/archnum/dolmen/internal/store"
	"github.com/archnum/sdk.application/container"
)

type (
	Config map[string]*store.Config

	configProvider interface {
		ConfigStores() Config
	}
)

func config(c container.Container) Config {
	return container.Value[configProvider](c, "config").ConfigStores()
}

/*
####### END ############################################################################################################
*/
