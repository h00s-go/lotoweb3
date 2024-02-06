package initializers

import (
	"github.com/go-raptor/raptor"
)

func App() *raptor.AppInitializer {
	services, controllers := ServicesAndControllers()

	return &raptor.AppInitializer{
		Services:    services,
		Controllers: controllers,
		Middlewares: Middlewares(),
	}
}
