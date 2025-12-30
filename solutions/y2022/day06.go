package y2022

import (
	"adventofcode/solutions/shared"
)

func Day06(input []string) (solution shared.Solution[int, int]) {

	solve := func(s string, n int) int {
		for i := n; i < len(s); i++ {
			seen := make(map[rune]bool)
			for _, r := range s[i-n : i] {
				seen[r] = true
			}
			if len(seen) == n {
				return i
			}
		}
		panic("could not solve for given input")
	}
	solution.Part1 = solve(input[0], 4)
	solution.Part2 = solve(input[0], 14)
	return
}
