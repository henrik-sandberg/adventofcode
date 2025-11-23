package y2024

import (
	"adventofcode/solutions/shared"
	"fmt"
)

func Day22(input []string) (solution shared.Solution[int, int]) {
	iterations := 2000
	iterate := func(i int) int {
		i ^= i << 6 & 0xFFFFFF
		i ^= i >> 5 & 0xFFFFFF
		i ^= i << 11 & 0xFFFFFF
		return i
	}
	for _, secret := range shared.IntSlice(input) {
		for i := 0; i < iterations; i++ {
			secret = iterate(secret)
		}
		solution.Part1 += secret
	}
	monkey := func(secret int) *map[string]int {
		windowLength := 4
		window := []int{}
		for i := 0; i < windowLength; i++ {
			prev := secret
			secret = iterate(secret)
			diff := secret%10 - prev%10
			window = append(window, diff)
		}
		prices := map[string]int{}
		for i := 0; i < iterations-windowLength; i++ {
			prev := secret
			secret = iterate(secret)
			diff := secret%10 - prev%10
			window = append(window[1:], diff)
			key := fmt.Sprint(window)
			if _, ok := prices[key]; !ok {
				prices[key] = secret % 10
			}
		}
		return &prices
	}
	prices := map[string]int{}
	for _, secret := range shared.IntSlice(input) {
		for k, v := range *monkey(secret) {
			prices[k] += v
		}
	}
	for _, v := range prices {
		if v > solution.Part2 {
			solution.Part2 = v
		}
	}
	return
}
