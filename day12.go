package main

import (
	"fmt"
	"strings"
)

func Day12(input []string) {
	grid := Map{strings.Join(input, ""), len(input[0])}.ToComplexGrid()
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
				// Check up
				if v == grid[p-1i] {
					if v != grid[p+1] && v != grid[p+1-1i] {
						longSides++
					}
					if v != grid[p-1] && v != grid[p-1-1i] {
						longSides++
					}
				}
				// Check right
				if v == grid[p+1] {
					if v != grid[p-1i] && v != grid[p+1-1i] {
						longSides++
					}
					if v != grid[p+1i] && v != grid[p+1+1i] {
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
	fmt.Println(part1)
	fmt.Println(part2)
}
