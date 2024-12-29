package y2024

import (
	"adventofcode/shared"
)

func Day10(input []string) shared.Solution {
	grid := shared.NewGrid(input)
	solver := func(distinctTrails bool) (res int) {
		for _, start := range grid.FindAll('0') {
			seen := map[complex64]bool{}
			queue := []complex64{start}
			var p complex64
			for len(queue) > 0 {
				p, queue = queue[0], queue[1:]
				for _, dir := range []complex64{1, -1, -1i, 1i} {
					if next := p + dir; (!seen[next] || distinctTrails) && grid[p]+1 == grid[next] {
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
	return shared.Solution{
		Part1: solver(false),
		Part2: solver(true),
	}
}
