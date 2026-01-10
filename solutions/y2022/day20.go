package y2022

import (
	"adventofcode/solutions/shared"
	"slices"
)

func Day20(input []string) (solution shared.Solution[int, int]) {
	solve := func(steps []int, rounds int) int {
		indices := make([]int, len(steps))
		for idx := range steps {
			indices[idx] = idx
		}
		for range rounds {
			for idx, step := range steps {
				i := slices.Index(indices, idx)

				copy(indices[i:], indices[i+1:])
				indices = indices[:len(indices)-1]

				newIndex := shared.PositiveMod(i+step, len(steps)-1)
				indices = append(indices, 0)
				copy(indices[newIndex+1:], indices[newIndex:])
				indices[newIndex] = idx
			}
		}
		zeroIdx := slices.Index(indices, slices.Index(steps, 0))
		sum := 0
		for _, n := range []int{1000, 2000, 3000} {
			sum += steps[indices[(zeroIdx+n)%len(indices)]]
		}
		return sum
	}
	steps := shared.IntSlice(input)
	solution.Part1 = solve(steps, 1)
	for idx := range steps {
		steps[idx] *= 811589153
	}
	solution.Part2 = solve(steps, 10)
	return
}
