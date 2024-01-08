package config

import "github.com/h00s/raptor"

func Routes() raptor.Routes {
	return raptor.CollectRoutes(
		raptor.Scope("/api/v1",
			raptor.Scope("/lotteries",
				raptor.Route("GET", "/6of45", "LotteriesController", "Get6of45"),
				raptor.Route("GET", "/7of35", "LotteriesController", "Get7of35"),
			),
		),
		raptor.Route("GET", "*", "SPAController", "Index"),
	)
}
