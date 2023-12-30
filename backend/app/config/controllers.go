package config

import (
	"github.com/h00s/lotoweb3/app/controllers"
	"github.com/h00s/raptor"
)

func Controllers() raptor.Controllers {
	lc := &controllers.LotteriesController{}
	sc := &controllers.SPAController{}

	return raptor.RegisterControllers(
		raptor.RegisterController(lc,
			lc.Get6of45,
			lc.Get7of39,
		),
		raptor.RegisterController(sc,
			sc.Index,
		),
	)
}
