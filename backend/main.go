package main

import (
	"github.com/go-raptor/raptor"
	"github.com/h00s/lotoweb3/config"
	"github.com/h00s/lotoweb3/config/initializers"
)

func main() {
	r := raptor.NewRaptor(
		initializers.App(),
		config.Routes(),
	)

	r.Listen()
}
