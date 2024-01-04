package config

import (
	"github.com/h00s/lotoweb3/app/controllers"
	"github.com/h00s/raptor"
)

func Controllers() raptor.Controllers {
	return raptor.RegisterControllers(
		&controllers.LotteriesController{
			Ls: Ls,
		},
		&controllers.SPAController{},
	)
}
