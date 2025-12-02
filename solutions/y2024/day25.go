package y2024

import (
	"adventofcode/solutions/shared"
)

func Day25(input []string) (solution shared.Solution[int, any]) {
	locks := [][]int{}
	keys := [][]int{}
	for r := 0; r < len(input); r += 8 {
		h := []int{}
		for col := range len(input[r]) {
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
	for _, key := range keys {
	lockloop:
		for _, lock := range locks {
			for i := range len(key) {
				if key[i]+lock[i] > 5 {
					continue lockloop
				}
			}
			solution.Part1++
		}
	}
	return
}
