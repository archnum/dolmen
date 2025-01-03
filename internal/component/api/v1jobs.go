/*
####### dolmen (c) 2024 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package api

import (
	"net/http"

	"github.com/archnum/sdk.base/uuid"
	"github.com/archnum/sdk.http/api"
	"github.com/archnum/sdk.http/api/bind"
	"github.com/archnum/sdk.http/api/failure"
	"github.com/archnum/sdk.http/api/render"
	"github.com/archnum/dolmen.jw/jw"
)

func (api *API) v1CreateJob(rr render.Renderer) error {
	jc := new(jw.JobCore)

	if err := bind.Body(rr, bind.DefaultMaxBodySize, jc); err != nil {
		return err
	}

	if jc.ID == "" {
		id, err := uuid.New()
		if err != nil {
			return failure.InternalServerError(err) ////////////////////////////////////////////////////////////////////
		}

		jc.ID = id
	}

	if err := jc.Validate(); err != nil {
		return failure.BadRequest(err) /////////////////////////////////////////////////////////////////////////////////
	}

	job, err := api.service.CreateJob(rr.Request().Context(), jc)
	if err != nil {
		return err
	}

	if job == nil {
		rr.NoContent()
		return nil
	}

	rr.WriteData(http.StatusCreated, job)

	return nil
}

func (api *API) v1Jobs(router api.Router) {
	router.Post("/", api.v1CreateJob)
}

/*
####### END ############################################################################################################
*/
