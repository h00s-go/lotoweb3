package components

import (
	"github.com/go-raptor/controllers/spa"
	"github.com/go-raptor/raptor/v3"
	"github.com/h00s/lotoweb3/app/controllers"
)

func Controllers() raptor.Controllers {
	return raptor.Controllers{
		&controllers.LotteriesController{},
		spa.NewSPAController("public", "index.html"),
	}
}
