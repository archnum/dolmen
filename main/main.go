/*
####### dolmen (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package main

import (
	"github.com/archnum/sdk.application/component/config"
	"github.com/archnum/sdk.application/component/crypto"
	"github.com/archnum/sdk.application/component/logger"
	"github.com/archnum/sdk.application/component/waitend"
	"github.com/archnum/sdk.application/container"
	"github.com/archnum/sdk.base/application"
	"github.com/archnum/sdk.http/component/server"

	"github.com/archnum/dolmen/internal/component/adapter"
	"github.com/archnum/dolmen/internal/component/api"
	"github.com/archnum/dolmen/internal/component/cmdline"
	"github.com/archnum/dolmen/internal/component/service"
	"github.com/archnum/dolmen/internal/component/stores"
	_cfg "github.com/archnum/dolmen/internal/config"
)

var (
	_version string
	_builtAt string
)

func main() {
	app, err := application.New(
		"dolmen",
		application.WithVersion(_version),
		application.WithBuiltAt(_builtAt),
	)
	if err == nil {
		c := container.New(app)

		c.AddComponents( ///////////////////// Liste ordonnée en fonction des dépendances //////////////////////////////
			crypto.New(c),
			config.New(c, new(_cfg.Config)),
			cmdline.New(c),
			logger.New(c),
			stores.New(c),
			adapter.New(c),
			service.New(c),
			api.New(c),
			server.New(c),
			waitend.New(c),
		)

		err = c.Run( //////////////////////////////// Liste par ordre d'exécution //////////////////////////////////////
			server.Name,
			waitend.Name,
		)
	}

	app.Exit(err)
}

/*
####### END ############################################################################################################
*/
