/*
####### dolmen (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package api

import (
	"github.com/archnum/sdk.application/component/logger"
	"github.com/archnum/sdk.application/container"
	"github.com/archnum/sdk.base/uuid"
	"github.com/archnum/sdk.http/api"

	"github.com/archnum/dolmen/internal/component/service"
)

type (
	implHandler struct {
		api.Manager
		service service.Service
	}
)

func newHandler(c container.Container) (*implHandler, error) {
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

	impl := &implHandler{
		Manager: api.New(p),
		service: service.Value(c),
	}

	impl.declareAPI()

	return impl, nil
}

func (impl *implHandler) declareAPI() {
	router := impl.Router()

	router.Mount("admin", impl.admin)
	router.Mount("/api/v1", impl.v1)
}

/*
####### END ############################################################################################################
*/
