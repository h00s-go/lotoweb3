package config

import (
	"github.com/go-raptor/raptor/v3/router"
)

func Routes() router.Routes {
	return router.CollectRoutes(
		router.Scope("api/v1",
			router.Scope("lotteries",
				router.Get("pick-one", "Lotteries.PickOne"),
				router.Get("pick-many", "Lotteries.PickMany"),
			),
		),
	)
}
