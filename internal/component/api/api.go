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

	"github.com/archnum/dolmen/internal/component/api/admin"
	v1 "github.com/archnum/dolmen/internal/component/api/v1"
)

type (
	implHandler struct {
		api.Manager
		admin *admin.API
		v1    *v1.API
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

	manager := api.New(p)

	admin, err := admin.New(c, manager)
	if err != nil {
		return nil, err
	}

	v1, err := v1.New(c, manager)
	if err != nil {
		return nil, err
	}

	impl := &implHandler{
		Manager: manager,
		admin:   admin,
		v1:      v1,
	}

	return impl, nil
}

/*
####### END ############################################################################################################
*/
