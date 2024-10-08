package config

import "github.com/go-raptor/raptor/v3"

func Routes() raptor.Routes {
	return raptor.CollectRoutes(
		raptor.Scope("/api/v1",
			raptor.Scope("/lotteries",
				raptor.Route("GET", "/pick-one", "LotteriesController", "PickOne"),
				raptor.Route("GET", "/pick-many", "LotteriesController", "PickMany"),
			),
		),
	)
}
