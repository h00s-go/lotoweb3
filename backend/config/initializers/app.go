package initializers

import (
	"github.com/h00s/lotoweb3/app/controllers"
	"github.com/h00s/lotoweb3/app/services"
	"github.com/h00s/raptor"
)

func App() *raptor.AppInitializer {
	ls := &services.LotteriesService{}

	return &raptor.AppInitializer{
		Services: raptor.Services{
			ls,
		},

		Middlewares: raptor.Middlewares{},

		Controllers: raptor.Controllers{
			&controllers.LotteriesController{
				Ls: ls,
			},
			&controllers.SPAController{},
		},
	}
}
