package y2025

import (
	"adventofcode/solutions/shared"
)

func Day03(input []string) (solution shared.Solution[int, int]) {
	solve := func(s string, length int) int {
		high := make([]byte, length)
		for start := range len(s) - length + 1 {
			for pt := range length {
				c := s[start+pt] - '0'
				if c > high[pt] {
					high[pt] = c
					clear(high[pt+1:])
				}
			}
		}
		result := 0
		for _, v := range high {
			result = result*10 + int(v)
		}
		return result
	}
	for _, line := range input {
		solution.Part1 += solve(line, 2)
		solution.Part2 += solve(line, 12)
	}
	return
}
