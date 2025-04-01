package main

import (
	"github.com/go-raptor/raptor/v3"
	"github.com/h00s/lotoweb3/app/utils"
	"github.com/h00s/lotoweb3/config"
	"github.com/h00s/lotoweb3/config/components"
)

func main() {
	app := raptor.New()

	logistiq, err := utils.NewLogistiqHandler(app.Utils.Config)
	if err == nil {
		app.Utils.SetHandler(logistiq)
		defer logistiq.Close()
	}

	app.Configure(components.New(app.Utils.Config))
	app.RegisterRoutes(config.Routes())
	app.Run()
}
