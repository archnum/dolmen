/*
####### dolmen (c) 2024 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package api

import (
	"github.com/archnum/sdk.http/api"
	"github.com/archnum/sdk.http/api/middleware"
)

func (api *API) v1(router api.Router) {

	router.Use(
		middleware.Logger(api.Logger()),
		middleware.Recover(api.Logger()),
	)

	router.Mount("/jobs", api.v1Jobs)
}

/*
####### END ############################################################################################################
*/
