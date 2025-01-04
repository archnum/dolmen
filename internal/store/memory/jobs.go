/*
####### dolmen (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package memory

import (
	"context"

	"github.com/archnum/sdk.base/failure"
	"github.com/archnum/sdk.base/kv"
	"github.com/archnum/sdk.jw/jw"
)

func (impl *implStore) CreateJob(ctx context.Context, job *jw.Job) (*jw.Job, error) {
	impl.mutex.Lock()
	defer impl.mutex.Unlock()

	_, ok := impl.jobs[job.ID]
	if ok {
		return nil,
			failure.New("A job with the same id already exists", kv.String("id", job.ID.String())) /////////////////////
	}

	if job.UniqueID != "" {
		for _, j := range impl.jobs {
			if j.UniqueID == job.UniqueID &&
				j.Status == jw.StatusTodo || j.Status == jw.StatusRunning || j.Status == jw.StatusPending {
				return nil, nil
			}
		}
	}

	impl.jobs[job.ID] = job

	return job, nil
}

/*
####### END ############################################################################################################
*/
