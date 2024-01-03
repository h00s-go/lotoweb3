package config

import (
	"github.com/h00s/lotoweb3/app/controllers"
	"github.com/h00s/lotoweb3/app/services"
	"github.com/h00s/raptor"
)

func Controllers() raptor.Controllers {
	ls := &services.LotteriesService{}

	return raptor.RegisterControllers(
		&controllers.LotteriesController{
			Ls: ls,
		},
		&controllers.SPAController{},
	)
}
