package main

import (
	"errors"
	"fmt"
	"strings"
)

type guard struct {
	loc, dir complex64
}

func Day06(input []string) {
	m := Map{strings.Join(input, ""), len(input[0])}
	grid := m.ToComplexGrid()
	start := guard{m.GetPointComplex(strings.Index(m.cells, "^")), -1i}
	path, _ := findPath(grid, start)
	unique := map[complex64]bool{}
	for k := range path {
		unique[k.loc] = true
	}
	fmt.Println(len(unique))
	part2 := 0
	for k, v := range grid {
		if v == '.' {
			grid[k] = '#'
			if _, err := findPath(grid, start); err != nil {
				part2++
			}
			grid[k] = '.'
		}
	}
	fmt.Println(part2)
}

func findPath(grid map[complex64]rune, g guard) (map[guard]bool, error) {
	visited := map[guard]bool{}
	for {
		visited[g] = true
		candidate := g.loc + g.dir
		if v, ok := grid[candidate]; !ok {
			return visited, nil
		} else if v == '#' {
			g.dir *= 1i
		} else {
			g.loc = candidate
		}
		if visited[g] {
			return nil, errors.New("Loop detected!")
		}
	}
}
