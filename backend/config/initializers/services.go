package initializers

import (
	"github.com/go-raptor/raptor/v3"
	"github.com/h00s/lotoweb3/app/services"
)

func Services() raptor.Services {
	return raptor.Services{
		&services.LotteriesService{},
	}
}
