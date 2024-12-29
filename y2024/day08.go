package y2024

import (
	"adventofcode/shared"
	"strings"
)

func Day08(input []string) (solution shared.Solution[int, int]) {
	grid := shared.NewGrid(input)
	frequencies := shared.RuneBag(strings.Join(input, ""))
	delete(frequencies, '.')
	solver := func(from, to int) (res int) {
		antinodes := map[complex64]bool{}
		for freq := range frequencies {
			towers := []complex64{}
			for k, v := range grid {
				if v == freq {
					towers = append(towers, k)
				}
			}
			for nums := range shared.Permutations(towers, 2) {
				a, b := nums[0], nums[1]
				for i := from; i < to; i++ {
					n := complex(float32(i), 0)
					antinodes[a+n*(a-b)] = true
				}
			}
		}
		for a := range antinodes {
			if _, ok := grid[a]; ok {
				res++
			}
		}
		return
	}
	solution.Part1 = solver(1, 2)
	solution.Part2 = solver(0, 50)
	return
}
