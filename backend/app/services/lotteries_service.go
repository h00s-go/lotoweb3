package services

import (
	"math/rand"
	"slices"

	"github.com/h00s/raptor"
)

type LotteriesService struct {
	raptor.Service
}

func (ls *LotteriesService) numbers(count, max int) []int {
	var numbers []int

	for i := 0; len(numbers) < count; i++ {
		number := rand.Intn(max) + 1
		if slices.Contains(numbers, number) {
			continue
		}
		numbers = append(numbers, number)
	}
	slices.Sort(numbers)
	return numbers
}

func (ls *LotteriesService) Get6of45() []int {
	return ls.numbers(6, 45)
}

func (ls *LotteriesService) Get7of39() []int {
	return ls.numbers(7, 39)
}
