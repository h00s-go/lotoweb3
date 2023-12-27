package config

import (
	"github.com/h00s/raptor"
)

func Routes() raptor.Routes {
	return raptor.Routes{
		raptor.Route("GET", "/api/v1/lotteries/6of45", "Lotteries", "Get6of45"),
		raptor.Route("GET", "/api/v1/lotteries/7of39", "Lotteries", "Get7of39"),
		raptor.Route("GET", "*", "SPA", "Index"),
	}
}
