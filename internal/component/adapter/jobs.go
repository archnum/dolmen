/*
####### dolmen (c) 2024 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package adapter

import (
	"context"

	"github.com/archnum/dolmen.jw/jw"
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
