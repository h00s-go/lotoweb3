package config

import (
	_ "embed"

	"github.com/go-raptor/raptor/v4/router"
)

//go:embed routes.yaml
var routesYAML []byte

func Routes() router.Routes {
	return router.CollectRoutes(
		//router.Get("/*path", "SPA.Index"),
		router.ParseYAML(routesYAML),
	)
}
