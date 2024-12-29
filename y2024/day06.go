package y2024

import (
	"adventofcode/shared"
	"errors"
)

type guard struct {
	location, direction complex64
}

func Day06(input []string) shared.Solution {
	grid := shared.NewGrid(input)
	start := guard{
		location:  grid.FindAny('^'),
		direction: -1i,
	}
	path, _ := findPath(grid, start)
	unique := make(map[complex64]bool)
	for k := range path {
		unique[k.location] = true
	}
	part1 := len(unique)
	delete(unique, start.location)
	part2 := 0
	for k := range unique {
		grid[k] = '#'
		if _, err := findPath(grid, start); err != nil {
			part2++
		}
		grid[k] = '.'
	}
	return shared.Solution{Part1: part1, Part2: part2}
}

func findPath(grid map[complex64]rune, g guard) (map[guard]bool, error) {
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
	return map[guard]bool{}, errors.New("Loop detected!")
}
