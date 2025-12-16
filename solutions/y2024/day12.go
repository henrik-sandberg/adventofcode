package y2024

import (
	"adventofcode/solutions/shared"
)

func Day12(input []string) (solution shared.Solution[int, int]) {
	grid := shared.NewGrid(input)
	seen := map[complex128]bool{}
	for point := range grid {
		if !seen[point] {
			seen[point] = true
			area := map[complex128]int{point: 0}
			queue := []complex128{point}
			for len(queue) > 0 {
				point, queue = queue[0], queue[1:]
				for _, dir := range []complex128{-1i, 1, 1i, -1} {
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
			solution.Part1 += len(area) * tmp
			solution.Part2 += len(area) * (tmp - longSides)
		}
	}
	return
}
