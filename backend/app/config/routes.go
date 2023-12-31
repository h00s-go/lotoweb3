package config

import (
	"github.com/h00s/raptor"
)

func Routes() raptor.Routes {
	return raptor.Routes{
		raptor.Route("GET", "/api/v1/lotteries/6of45", "LotteriesController", "Get6of45"),
		raptor.Route("GET", "/api/v1/lotteries/7of39", "LotteriesController", "Get7of39"),
		raptor.Route("GET", "*", "SPAController", "Index"),
	}
}
