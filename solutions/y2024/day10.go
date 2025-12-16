package y2024

import (
	"adventofcode/solutions/shared"
)

func Day10(input []string) (solution shared.Solution[int, int]) {
	grid := shared.NewGrid(input)
	solver := func(distinctTrails bool) (res int) {
		for _, point := range grid.FindAll('0') {
			seen := map[complex128]bool{}
			queue := []complex128{point}
			for len(queue) > 0 {
				point, queue = queue[0], queue[1:]
				for _, dir := range []complex128{-1i, 1, 1i, -1} {
					if next := point + dir; (!seen[next] || distinctTrails) && grid[point]+1 == grid[next] {
						seen[next] = true
						if grid[next] == '9' {
							res++
						} else {
							queue = append(queue, next)
						}
					}
				}
			}
		}
		return
	}
	solution.Part1 = solver(false)
	solution.Part2 = solver(true)
	return
}
