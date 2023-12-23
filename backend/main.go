package main

import (
	"github.com/h00s/lotoweb3/app/config"
	"github.com/h00s/raptor"
)

func main() {
	r := raptor.NewRaptorAPI(raptor.Config{
		Address: "localhost",
		Port:    3000,
	})

	r.Controllers(config.Controllers())
	r.Routes(config.Routes())

	r.Listen()
}
