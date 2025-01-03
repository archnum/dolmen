/*
####### dolmen (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package mongodb

import (
	"context"

	"github.com/archnum/dolmen.jw/jw"
	"github.com/archnum/sdk.base/failure"
)

func (impl *implStore) CreateJob(ctx context.Context, job *jw.Job) (*jw.Job, error) {
	return nil, failure.NotImplemented
}

/*
####### END ############################################################################################################
*/
