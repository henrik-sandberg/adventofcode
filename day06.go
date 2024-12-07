package main

import (
	"errors"
	"fmt"
	"strings"
)

type guard struct {
	location  int
	direction int
}

func (g guard) move(m Map) (guard, error) {
	candidate := -1
	if g.direction == 0 && g.location-m.Width() >= 0 {
		candidate = g.location - m.Width()
	} else if g.direction == 1 && g.location%m.Width()+1 < m.Width() {
		candidate = g.location + 1
	} else if g.direction == 2 && g.location+m.Width() < len(m.cells) {
		candidate = g.location + m.Width()
	} else if g.direction == 3 && g.location%m.Width()-1 >= 0 {
		candidate = g.location - 1
	}
	if candidate == -1 {
		return guard{}, errors.New("Run out of map")
	} else if m.cells[candidate] != '#' {
		return guard{candidate, g.direction}, nil
	} else {
		return guard{g.location, (g.direction + 1) % 4}, nil
	}
}

func Day06(input []string) {
	m := Map{strings.Join(input, ""), len(input[0])}
	part1, _ := findPath(m)
	part2 := 0
	for i := range m.cells {
		if m.cells[i] == '.' {
			runes := []rune(m.cells)
			runes[i] = '#'
			new_map := Map{string(runes), m.Width()}
			_, err := findPath(new_map)
			if err != nil {
				part2++
			}
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

func findPath(m Map) (int, error) {
	visited := make(map[guard]bool)
	g := guard{strings.Index(m.cells, "^"), 0}
	var runOutOfMapError error
	for runOutOfMapError == nil && !visited[g] {
		visited[g] = true
		g, runOutOfMapError = g.move(m)
	}
	if runOutOfMapError == nil {
		return 0, errors.New("Stuck in loop")
	}
	unique := map[int]bool{}
	for k := range visited {
		unique[k.location] = true
	}
	return len(unique), nil
}
