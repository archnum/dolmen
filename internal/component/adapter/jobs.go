/*
####### dolmen (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package adapter

import (
	"context"

	"github.com/archnum/sdk.jw/jw"
)

func (impl *implComponent) CreateJob(ctx context.Context, job *jw.Job) (*jw.Job, error) {
	job, err := impl.store.CreateJob(ctx, job)
	if err != nil {
		return nil, err
	}

	return job, nil
}

/*
####### END ############################################################################################################
*/
