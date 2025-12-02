package y2024

import (
	"fmt"

	"adventofcode/solutions/shared"
)

type guard struct {
	location, direction complex128
}

func Day06(input []string) (solution shared.Solution[int, int]) {
	grid := shared.NewGrid(input)
	start := guard{
		location:  grid.FindAny('^'),
		direction: -1i,
	}
	path, _ := findPath(grid, start)
	unique := make(map[complex128]bool)
	for k := range path {
		unique[k.location] = true
	}
	solution.Part1 = len(unique)
	delete(unique, start.location)
	for k := range unique {
		grid[k] = '#'
		if _, err := findPath(grid, start); err != nil {
			solution.Part2++
		}
		grid[k] = '.'
	}
	return
}

func findPath(grid map[complex128]byte, g guard) (map[guard]bool, error) {
	visited := make(map[guard]bool, len(grid))
	for !visited[g] {
		visited[g] = true
		next := g.location + g.direction
		if v, ok := grid[next]; !ok {
			return visited, nil
		} else if v == '#' {
			g.direction *= 1i
		} else {
			g.location = next
		}
	}
	return nil, fmt.Errorf("loop detected")
}
