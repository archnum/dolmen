/*
####### dolmen (c) 2024 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package mongodb

import (
	"context"

	"github.com/archnum/sdk.base/failure"
	"github.com/archnum/dolmen.jw/jw"
)

func (impl *implStore) CreateJob(ctx context.Context, job *jw.Job) (*jw.Job, error) {
	return nil, failure.NotImplemented
}

/*
####### END ############################################################################################################
*/
