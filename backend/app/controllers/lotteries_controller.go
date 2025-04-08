package controllers

import (
	"strconv"

	"github.com/go-raptor/raptor/v3"
	"github.com/h00s/lotoweb3/app/services"
)

type LotteriesController struct {
	raptor.Controller

	Lotteries *services.LotteriesService
}

func (hc *LotteriesController) PickOne(s raptor.State) error {
	if s.QueryParam("numbers") != "" && s.QueryParam("max") != "" {
		numbers, err := strconv.Atoi(s.QueryParam("numbers"))
		if err != nil {
			return err
		}
		max, err := strconv.Atoi(s.QueryParam("max"))
		if err != nil {
			return err
		}
		return s.JSONResponse(hc.Lotteries.PickOne(numbers, max))
	}
	return s.JSONResponse(hc.Lotteries.PickOne(6, 45))
}

func (hc *LotteriesController) PickMany(s raptor.State) error {
	count, err := strconv.Atoi(s.QueryParam("count"))
	if err != nil {
		return err
	}

	numbers, err := strconv.Atoi(s.QueryParam("numbers"))
	if err != nil {
		return err
	}

	max, err := strconv.Atoi(s.QueryParam("max"))
	if err != nil {
		return err
	}

	return s.JSONResponse(hc.Lotteries.PickMany(count, numbers, max))
}
