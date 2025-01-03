/*
####### dolmen (c) 2024 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package api

import (
	"github.com/archnum/sdk.http/api"
	"github.com/archnum/sdk.http/api/middleware"
	"github.com/archnum/sdk.http/api/render"
	_loggers "github.com/archnum/sdk.loggers"
	_external "github.com/ltrochet/loggers"
)

func (api *API) admin(router api.Router) {
	router.Use(
		middleware.Logger(api.Logger()),
		middleware.Recover(api.Logger()),
	)

	router.Get(
		"/loggers",
		func(rr render.Renderer) error {
			_external.Handler(_loggers.All()).ServeHTTP(rr.ResponseWriter(), rr.Request())
			return nil
		},
	)

	// chi.middleware.Profiler() ?
}

/*
####### END ############################################################################################################
*/