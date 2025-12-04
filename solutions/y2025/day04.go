package y2025

import (
	"adventofcode/solutions/shared"
	"maps"
)

func Day04(input []string) (solution shared.Solution[int, int]) {
	adjecent := []complex128{
		1, -1, 1i, -1i, 1 + 1i, 1 - 1i, -1 + 1i, -1 - 1i,
	}
	grid := shared.NewGrid(input)
	maps.DeleteFunc(grid, func(_ complex128, v byte) bool {
		return v != '@'
	})
	countNeighbours := func(c complex128) int {
		result := 0
		for _, a := range adjecent {
			if grid[c+a] == '@' {
				result++
			}
		}
		return result
	}
	for k := range grid {
		if countNeighbours(k) < 4 {
			solution.Part1++
		}
	}
	original := len(grid)
	for removed := true; removed; {
		removed = false
		for k := range grid {
			if countNeighbours(k) < 4 {
				delete(grid, k)
				removed = true
			}
		}
	}
	solution.Part2 = original - len(grid)
	return
}
