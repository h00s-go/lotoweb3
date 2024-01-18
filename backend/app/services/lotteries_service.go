package services

import (
	"crypto/rand"
	"math/big"
	"slices"

	"github.com/go-raptor/raptor"
)

type LotteriesService struct {
	raptor.Service
}

func (ls *LotteriesService) numbers(count, max int) []int {
	var numbers []int

	for i := 0; len(numbers) < count; i++ {
		numberBig, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
		if err != nil {
			panic(err)
		}
		number := int(numberBig.Int64()) + 1
		if slices.Contains(numbers, number) {
			continue
		}
		numbers = append(numbers, number)
	}
	slices.Sort(numbers)
	return numbers
}

func (ls *LotteriesService) PickOne(numbers, max int) []int {
	return ls.numbers(numbers, max)
}

func (ls *LotteriesService) PickMany(count, numbers, max int) [][]int {
	collection := make([][]int, count)
	for i := 0; i < count; i++ {
		collection[i] = ls.numbers(numbers, max)
	}
	return collection
}
