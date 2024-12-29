package y2024

import (
	"adventofcode/shared"
)

func Day25(input []string) shared.Solution {
	locks := [][]int{}
	keys := [][]int{}
	for r := 0; r < len(input); r += 8 {
		h := []int{}
		for col := 0; col < len(input[r]); col++ {
			count := 0
			for i := 1; i < 6; i++ {
				if input[r+i][col] == '#' {
					count++
				}
			}
			h = append(h, count)
		}
		if input[r][0] == '#' {
			locks = append(locks, h)
		} else {
			keys = append(keys, h)
		}
	}
	part1 := 0
	for _, key := range keys {
	lockloop:
		for _, lock := range locks {
			for i := 0; i < len(key); i++ {
				if key[i]+lock[i] > 5 {
					continue lockloop
				}
			}
			part1++
		}
	}
	return shared.Solution{Part1: part1, Part2: nil}
}
