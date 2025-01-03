/*
####### dolmen (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package stores

import (
	"context"

	"github.com/archnum/dolmen.jw/jw"
	"github.com/archnum/sdk.base/failure"
	"github.com/archnum/sdk.base/kv"
)

func (impl *implComponent) CreateJob(ctx context.Context, job *jw.Job) (*jw.Job, error) {
	store, ok := impl.namespaces[job.Namespace]
	if !ok {
		return nil,
			failure.New("unknown namespace", kv.String("namespace", job.Namespace)) ////////////////////////////////////
	}

	return store.CreateJob(ctx, job)
}

/*
####### END ############################################################################################################
*/
