package config

import (
	"github.com/h00s/lotoweb3/app/controllers"
	"github.com/h00s/raptor"
)

func Controllers() raptor.Controllers {
	lc := &controllers.LotteriesController{}
	sc := &controllers.SPAController{}

	return raptor.RegisterControllers(
		raptor.RegisterController("Lotteries", &lc.Controller,
			raptor.Action("Get6of45", lc.Get6of45),
			raptor.Action("Get7of39", lc.Get7of39),
		),
		raptor.RegisterController("SPA", &sc.Controller,
			raptor.Action("Index", sc.Index),
		),
	)
}
