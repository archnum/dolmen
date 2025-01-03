/*
####### dolmen (c) 2024 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package stores

import (
	"errors"

	"github.com/archnum/sdk.application/component/logger"
	"github.com/archnum/sdk.application/container"
	"github.com/archnum/sdk.base/failure"
	"github.com/archnum/sdk.base/kv"

	"github.com/archnum/dolmen/internal/store"
)

const (
	_name = "stores"
)

type (
	implComponent struct {
		*container.Component
		stores     map[string]store.Store
		namespaces map[string]store.Store
	}
)

func New(c container.Container) *implComponent {
	return &implComponent{
		Component:  container.NewComponent(_name, c),
		stores:     make(map[string]store.Store),
		namespaces: make(map[string]store.Store),
	}
}

func Value(c container.Container) store.Store {
	return container.Value[store.Store](c, _name)
}

//////////////////////
/// Implementation ///
//////////////////////

func (impl *implComponent) Build() error {
	c := impl.C()
	logger := logger.Value(c)
	cfg := config(c)

	for name, sCfg := range cfg {
		if sCfg.Disabled {
			continue
		}

		if len(sCfg.Namespaces) == 0 {
			continue
		}

		if sCfg.Type == "" {
			sCfg.Type = name
		}

		store, err := store.Build(name, sCfg)
		if err != nil {
			_ = impl.Close()
			return err
		}

		if _, ok := impl.stores[name]; ok {
			_ = impl.Close()
			return failure.New( ////////////////////////////////////////////////////////////////////////////////////////
				"this store name is used more than once",
				kv.String("name", name),
			)
		}

		impl.stores[name] = store

		logger.Info("Store", kv.String("name", name), kv.String("type", sCfg.Type)) //::::::::::::::::::::::::::::::::::

		for _, namespace := range sCfg.Namespaces {
			if _, ok := impl.namespaces[namespace]; ok {
				_ = impl.Close()
				return failure.New( ////////////////////////////////////////////////////////////////////////////////////
					"this namespace is used more than once",
					kv.String("name", namespace),
				)
			}

			impl.namespaces[namespace] = store

			logger.Info("Namespace", kv.String("name", namespace), kv.String("store", name)) //:::::::::::::::::::::::::
		}
	}

	impl.SetValue(impl)

	return nil
}

func (impl *implComponent) Close() error {
	var errs error

	for name, store := range impl.stores {
		if err := store.Close(); err != nil {
			errs = errors.Join(
				errs,
				failure.WithMessage(err, "failed to close this store", kv.String("name", name)), ///////////////////////
			)
		}
	}

	return errs
}

/*
####### END ############################################################################################################
*/
