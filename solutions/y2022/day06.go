package y2022

import (
	"adventofcode/solutions/shared"
)

func Day06(input []string) (solution shared.Solution[int, int]) {
	solution.Part1 = findIndexOfDistinctCharacters(input[0], 4)
	solution.Part2 = findIndexOfDistinctCharacters(input[0], 14)
	return
}

// Find index where the subsequent characters are all distinct
func findIndexOfDistinctCharacters(s string, n int) int {
	for i := n; i < len(s); i++ {
		unique := map[rune]bool{}
		for _, r := range s[i-n : i] {
			unique[r] = true
		}
		if len(unique) == n {
			return i
		}
	}
	return 0
}
