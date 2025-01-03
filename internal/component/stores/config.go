/*
####### dolmen (c) 2024 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package stores

import (
	"github.com/archnum/sdk.application/container"
	"github.com/archnum/dolmen/internal/store"
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
