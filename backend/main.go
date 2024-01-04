package main

import (
	"github.com/h00s/lotoweb3/app/config"
	"github.com/h00s/raptor"
)

func main() {
	r := raptor.NewRaptor()

	r.Services(config.Services())
	r.Controllers(config.Controllers())
	r.Routes(config.Routes())

	r.Listen()
}
