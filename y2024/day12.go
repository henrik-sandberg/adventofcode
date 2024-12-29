package y2024

import (
	"adventofcode/shared"
)

func Day12(input []string) shared.Solution {
	grid := shared.NewGrid(input)
	part1 := 0
	part2 := 0
	seen := map[complex64]bool{}
	for point := range grid {
		if !seen[point] {
			seen[point] = true
			area := map[complex64]int{point: 0}
			queue := []complex64{point}
			for len(queue) > 0 {
				point, queue = queue[0], queue[1:]
				for _, dir := range []complex64{1, -1, -1i, 1i} {
					if next := point + dir; grid[next] != grid[point] {
						area[point]++
					} else if !seen[next] {
						seen[next] = true
						area[next] = 0
						queue = append(queue, next)
					}
				}
			}
			longSides := 0
			for p := range area {
				v := grid[p]
				if v == grid[p-1i] {
					if right := p + 1; v != grid[right] && v != grid[right-1i] {
						longSides++
					}
					if left := p - 1; v != grid[left] && v != grid[left-1i] {
						longSides++
					}
				}
				if v == grid[p+1] {
					if up := p - 1i; v != grid[up] && v != grid[up+1] {
						longSides++
					}
					if down := p + 1i; v != grid[down] && v != grid[down+1] {
						longSides++
					}
				}
			}
			tmp := 0
			for _, v := range area {
				tmp += v
			}
			part1 += len(area) * tmp
			part2 += len(area) * (tmp - longSides)
		}
	}
	return shared.Solution{Part1: part1, Part2: part2}
}
