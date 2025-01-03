/*
####### dolmen (c) 2024 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package adapter

import (
	"context"

	"github.com/archnum/sdk.application/container"
	"github.com/archnum/dolmen.jw/jw"

	"github.com/archnum/dolmen/internal/component/stores"
	"github.com/archnum/dolmen/internal/store"
)

const (
	_name = "adapter"
)

type (
	Adapter interface {
		CreateJob(ctx context.Context, job *jw.Job) (*jw.Job, error)
	}

	implComponent struct {
		*container.Component
		store store.Store
	}
)

func New(c container.Container) *implComponent {
	return &implComponent{
		Component: container.NewComponent(_name, c),
	}
}

func Value(c container.Container) Adapter {
	return container.Value[Adapter](c, _name)
}

//////////////////////
/// Implementation ///
//////////////////////

func (impl *implComponent) Build() error {
	c := impl.C()
	impl.store = stores.Value(c)

	impl.SetValue(impl)

	return nil
}

/*
####### END ############################################################################################################
*/
