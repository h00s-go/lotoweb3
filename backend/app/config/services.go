package config

import (
	"github.com/h00s/lotoweb3/app/services"
	"github.com/h00s/raptor"
)

var (
	Ls = &services.LotteriesService{}
)

func Services() raptor.Services {
	return raptor.Services{
		Ls,
	}
}
