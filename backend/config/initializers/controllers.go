package initializers

import (
	"github.com/go-raptor/raptor/v3"
	"github.com/h00s/lotoweb3/app/controllers"
)

func Controllers() raptor.Controllers {
	return raptor.Controllers{
		&controllers.LotteriesController{},
	}
}
