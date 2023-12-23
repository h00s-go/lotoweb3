package controllers

import (
	"github.com/h00s/lotoweb3/app/models"
	"github.com/h00s/raptor"
)

type LotteriesController struct {
	raptor.Controller
}

func (hc *LotteriesController) Get6of45(c *raptor.Context) error {
	return c.JSON(models.Numbers(6, 45))
}

func (hc *LotteriesController) Get7of39(c *raptor.Context) error {
	return c.JSON(models.Numbers(7, 39))
}
