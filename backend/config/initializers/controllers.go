package initializers

import (
	"github.com/go-raptor/raptor"
	"github.com/h00s/lotoweb3/app/controllers"
	"github.com/h00s/lotoweb3/app/services"
)

func ServicesAndControllers() (raptor.Services, raptor.Controllers) {
	ls := &services.LotteriesService{}

	return raptor.Services{
			ls,
		},
		raptor.Controllers{
			&controllers.LotteriesController{
				Ls: ls,
			},
			&controllers.SPAController{},
		}
}
