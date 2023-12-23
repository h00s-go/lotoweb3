package models

import (
	"math/rand"
	"slices"
)

func Numbers(count, max int) []int {
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
