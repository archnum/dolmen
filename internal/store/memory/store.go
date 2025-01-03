/*
####### dolmen (c) 2024 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package memory

import (
	"sync"

	"github.com/archnum/sdk.base/uuid"
	"github.com/archnum/dolmen.jw/jw"
)

type (
	implStore struct {
		name  string
		jobs  map[uuid.UUID]*jw.Job
		mutex sync.Mutex
	}
)

func New(name string, cfg map[string]any) (*implStore, error) {
	impl := &implStore{
		name: name,
		jobs: make(map[uuid.UUID]*jw.Job),
	}

	return impl, nil
}

func (impl *implStore) Close() error {
	return nil
}

/*
####### END ############################################################################################################
*/
