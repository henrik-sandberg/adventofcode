package y2024

import (
	"adventofcode/shared"
	"fmt"
)

func Day22(input []string) shared.Solution {
	iterations := 2000
	iterate := func(i int) int {
		i ^= i << 6 & 0xFFFFFF
		i ^= i >> 5 & 0xFFFFFF
		i ^= i << 11 & 0xFFFFFF
		return i
	}
	part1 := 0
	for _, secret := range shared.IntSlice(input) {
		for i := 0; i < iterations; i++ {
			secret = iterate(secret)
		}
		part1 += secret
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
	part2 := 0
	for _, v := range prices {
		if v > part2 {
			part2 = v
		}
	}
	return shared.Solution{Part1: part1, Part2: part2}
}
