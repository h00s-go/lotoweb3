package initializers

import (
	"github.com/go-raptor/raptor"
)

func App(c *raptor.Config) *raptor.AppInitializer {
	return &raptor.AppInitializer{
		Services:    Services(),
		Middlewares: Middlewares(),
		Controllers: Controllers(),
	}
}
