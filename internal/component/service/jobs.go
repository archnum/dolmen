/*
####### dolmen (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package service

import (
	"context"

	"github.com/archnum/sdk.base/logger/level"
	"github.com/archnum/sdk.jw/jw"
)

func (impl *implComponent) CreateJob(ctx context.Context, jc *jw.JobCore) (*jw.Job, error) {
	job, err := impl.adapter.CreateJob(ctx, jc.NewJob())
	if err != nil {
		jc.LogError(err, impl.logger, "Failed to create this job") //:::::::::::::::::::::::::::::::::::::::::::::::::::
		return nil, err
	}

	if job == nil {
		jc.Log(impl.logger, level.Notice, "A job with the same unique id already exists") //::::::::::::::::::::::::::::
		return nil, nil
	}

	job.Log(impl.logger, level.Info, "New job") //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

	return job, nil
}

/*
####### END ############################################################################################################
*/
