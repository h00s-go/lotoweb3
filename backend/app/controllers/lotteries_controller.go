package controllers

import (
	"github.com/h00s/lotoweb3/app/services"
	"github.com/h00s/raptor"
)

type LotteriesController struct {
	raptor.Controller

	Ls *services.LotteriesService
}

func (hc *LotteriesController) Get6of45(c *raptor.Context) error {
	return c.JSON(hc.Ls.Get6of45())
}

func (hc *LotteriesController) Get7of35(c *raptor.Context) error {
	return c.JSON(hc.Ls.Get7of35())
}
