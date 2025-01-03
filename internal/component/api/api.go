/*
####### dolmen (c) 2024 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package api

import (
	"net/http"

	"github.com/archnum/sdk.application/component/logger"
	"github.com/archnum/sdk.application/container"
	"github.com/archnum/sdk.base/uuid"
	"github.com/archnum/sdk.http/api"

	"github.com/archnum/dolmen/internal/component/service"
)

type (
	API struct {
		api.Manager
		service service.Service
	}
)

func newAPI(c container.Container) (*API, error) {
	logger := logger.Value(c)

	id, err := uuid.New()
	if err != nil {
		return nil, err
	}

	logger = logger.New(id, "api")
	logger.Register()

	p := &api.Params{
		Logger: logger,
	}

	api := &API{
		Manager: api.New(p),
		service: service.Value(c),
	}

	return api, nil
}

func newHandler(c container.Container) (http.Handler, error) {
	api, err := newAPI(c)
	if err != nil {
		return nil, err
	}

	router := api.Router()

	router.Mount("admin", api.admin)
	router.Mount("/api/v1", api.v1)

	return api, nil
}

/*
####### END ############################################################################################################
*/
