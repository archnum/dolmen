/*
####### dolmen (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package v1

import (
	"github.com/archnum/sdk.application/container"
	"github.com/archnum/sdk.http/api"
	"github.com/archnum/sdk.http/api/middleware"

	"github.com/archnum/dolmen/internal/component/service"
)

type (
	API struct {
		api.Manager
		service service.Service
	}
)

func New(c container.Container, manager api.Manager) (*API, error) {
	api := &API{
		Manager: manager,
		service: service.Value(c),
	}

	manager.Router().Mount("/api/v1", api.v1)

	return api, nil
}

func (api *API) v1(router api.Router) {

	router.Use(
		middleware.Logger(api.Logger()),
		middleware.Recover(api.Logger()),
	)

	router.Mount("/jobs", api.jobs)
}

/*
####### END ############################################################################################################
*/
