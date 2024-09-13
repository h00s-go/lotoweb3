package initializers

import (
	"github.com/go-raptor/raptor/v3"
	"github.com/h00s/lotoweb3/config"
)

func App(c *raptor.Config) *raptor.AppInitializer {
	return &raptor.AppInitializer{
		Routes:      config.Routes(),
		Services:    Services(),
		Middlewares: Middlewares(),
		Controllers: Controllers(),
	}
}
