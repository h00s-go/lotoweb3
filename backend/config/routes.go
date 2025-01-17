package config

import "github.com/go-raptor/raptor/v3"

func Routes() raptor.Routes {
	return raptor.CollectRoutes(
		raptor.Scope("api/v1",
			raptor.Scope("lotteries",
				raptor.Get("pick-one", "Lotteries#PickOne"),
				raptor.Get("pick-many", "Lotteries#PickMany"),
			),
		),
	)
}
