package controllers

import (
	"strconv"

	"github.com/go-raptor/raptor/v2"
	"github.com/h00s/lotoweb3/app/services"
)

type LotteriesController struct {
	raptor.Controller

	Lotteries *services.LotteriesService
}

func (hc *LotteriesController) PickOne(c *raptor.Context) error {
	if c.Query("numbers") != "" && c.Query("max") != "" {
		numbers, err := strconv.Atoi(c.Query("numbers"))
		if err != nil {
			return err
		}
		max, err := strconv.Atoi(c.Query("max"))
		if err != nil {
			return err
		}
		return c.JSON(hc.Lotteries.PickOne(numbers, max))
	}
	return c.JSON(hc.Lotteries.PickOne(6, 45))
}

func (hc *LotteriesController) PickMany(c *raptor.Context) error {
	count, err := strconv.Atoi(c.Query("count"))
	if err != nil {
		return err
	}

	numbers, err := strconv.Atoi(c.Query("numbers"))
	if err != nil {
		return err
	}

	max, err := strconv.Atoi(c.Query("max"))
	if err != nil {
		return err
	}

	return c.JSON(hc.Lotteries.PickMany(count, numbers, max))
}
