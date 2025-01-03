/*
####### dolmen (c) 2024 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package store

import (
	"context"

	"github.com/archnum/sdk.base/failure"
	"github.com/archnum/sdk.base/kv"
	"github.com/archnum/dolmen.jw/jw"

	"github.com/archnum/dolmen/internal/store/memory"
	"github.com/archnum/dolmen/internal/store/mongodb"
)

type (
	Store interface {
		CreateJob(ctx context.Context, job *jw.Job) (*jw.Job, error)
		Close() error
	}
)

func Build(name string, cfg *Config) (Store, error) {
	switch cfg.Type {
	case "memory":
		return memory.New(name, cfg.Config)
	case "mongodb":
		return mongodb.New(name, cfg.Config)
	default:
		return nil, failure.New("unknown store", kv.String("name", name)) //////////////////////////////////////////////
	}
}

/*
####### END ############################################################################################################
*/
