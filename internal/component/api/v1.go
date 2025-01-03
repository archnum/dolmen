/*
####### dolmen (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package api

import (
	"github.com/archnum/sdk.http/api"
	"github.com/archnum/sdk.http/api/middleware"
)

func (impl *implHandler) v1(router api.Router) {

	router.Use(
		middleware.Logger(impl.Logger()),
		middleware.Recover(impl.Logger()),
	)

	router.Mount("/jobs", impl.v1Jobs)
}

/*
####### END ############################################################################################################
*/
