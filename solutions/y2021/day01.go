package y2021

import (
	"adventofcode/solutions/shared"
)

func Day01(input []string) (solution shared.Solution[int, int]) {
	depths := shared.IntSlice(input)
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			solution.Part1++
		}
	}
	for i := 3; i < len(depths); i++ {
		if depths[i] > depths[i-3] {
			solution.Part2++
		}
	}
	return
}
