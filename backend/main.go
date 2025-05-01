package main

import (
	"github.com/go-raptor/raptor/v4"
	"github.com/h00s/lotoweb3/app/utils"
	"github.com/h00s/lotoweb3/config"
	"github.com/h00s/lotoweb3/config/components"
)

func main() {
	app := raptor.New()

	logistiq, err := utils.NewLogistiqHandler(app.Core.Resources.Config)
	if err == nil {
		app.Core.Resources.SetLogHandler(logistiq)
		defer logistiq.Close()
	}

	app.Configure(components.New(app.Core.Resources.Config))
	app.RegisterRoutes(config.Routes())
	app.Run()
}
