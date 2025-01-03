/*
####### dolmen (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package service

import (
	"context"

	"github.com/archnum/dolmen.jw/jw"
	_logger "github.com/archnum/sdk.application/component/logger"
	"github.com/archnum/sdk.application/container"
	"github.com/archnum/sdk.base/logger"

	"github.com/archnum/dolmen/internal/component/adapter"
)

const (
	_name = "service"
)

type (
	Service interface {
		CreateJob(ctx context.Context, jc *jw.JobCore) (*jw.Job, error)
	}

	implComponent struct {
		*container.Component
		logger  *logger.Logger
		adapter adapter.Adapter
	}
)

func New(c container.Container) *implComponent {
	return &implComponent{
		Component: container.NewComponent(_name, c),
	}
}

func Value(c container.Container) Service {
	return container.Value[Service](c, _name)
}

//////////////////////
/// Implementation ///
//////////////////////

func (impl *implComponent) Build() error {
	c := impl.C()
	impl.logger = _logger.Value(c)
	impl.adapter = adapter.Value(c)

	impl.SetValue(impl)

	return nil
}

/*
####### END ############################################################################################################
*/
