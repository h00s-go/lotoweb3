package controllers

import (
	"github.com/go-raptor/raptor"
	"github.com/h00s/lotoweb3/app/services"
)

type LotteriesController struct {
	raptor.Controller

	Ls *services.LotteriesService
}

func (hc *LotteriesController) PickOne(c *raptor.Context) error {
	numbers := c.QueryInt("numbers", 6)
	max := c.QueryInt("max", 45)
	return c.JSON(hc.Ls.PickOne(numbers, max))
}

func (hc *LotteriesController) PickMany(c *raptor.Context) error {
	count := c.QueryInt("count", 5)
	numbers := c.QueryInt("numbers", 6)
	max := c.QueryInt("max", 45)
	return c.JSON(hc.Ls.PickMany(count, numbers, max))
}
