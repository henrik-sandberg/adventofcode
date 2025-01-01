package y2023

import (
	"adventofcode/shared"
	"strings"
)

func Day04(input []string) (solution shared.Solution[int, int]) {
	part2 := make([]int, len(input))
	for i := range part2 {
		part2[i] = 1
	}
	for i, row := range input {
		cards := strings.Split(strings.SplitN(row, ": ", 2)[1], " | ")
		wins := len(shared.Intersect(strings.Fields(cards[0]), strings.Fields(cards[1])))
		if wins > 0 {
			solution.Part1 += 1 << (wins - 1)
		}
		for j := i + 1; j < i+wins+1; j++ {
			part2[j] += part2[i]
		}
	}
	solution.Part2 = shared.Sum(part2...)
	return
}
