package initializers

import (
	"github.com/go-raptor/raptor"
	"github.com/h00s/lotoweb3/app/controllers"
)

func Controllers() raptor.Controllers {
	return raptor.Controllers{
		&controllers.LotteriesController{},
		&controllers.SPAController{},
	}
}
