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

func (g guard) move(m string, width int) (guard, error) {
	candidate := -1
	if g.direction == 0 && g.location-width >= 0 {
		candidate = g.location - width
	} else if g.direction == 1 && g.location%width+1 < width {
		candidate = g.location + 1
	} else if g.direction == 2 && g.location+width < len(m) {
		candidate = g.location + width
	} else if g.direction == 3 && g.location%width-1 >= 0 {
		candidate = g.location - 1
	}
	if candidate == -1 {
		return guard{}, errors.New("Run out of map")
	} else if m[candidate] != '#' {
		return guard{candidate, g.direction}, nil
	} else {
		return guard{g.location, (g.direction + 1) % 4}, nil
	}
}

func Day06(input []string) {
	width, height := len(input[0]), len(input)
	fmt.Printf("width: %d, height: %d\n", width, height)
	m := strings.Join(input, "")
	part1, _ := findPath(m, width)
	part2 := 0
	for i := range m {
		if m[i] == '.' {
			runes := []rune(m)
			runes[i] = '#'
			new_map := string(runes)
			_, err := findPath(new_map, width)
			if err != nil {
				part2++
			}
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

func findPath(m string, width int) (int, error) {
	visited := make(map[guard]bool)
	g := guard{strings.Index(m, "^"), 0}
	var runOutOfMapError error
	for runOutOfMapError == nil && !visited[g] {
		visited[g] = true
		g, runOutOfMapError = g.move(m, width)
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
