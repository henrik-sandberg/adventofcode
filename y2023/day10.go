package y2023

import (
	"adventofcode/shared"
	"slices"
	"strings"
)

func Day10(input []string) (solution shared.Solution[int, int]) {
	grid := shared.NewGrid(input)
	findLoop := func(start complex64, direction complex64) []complex64 {
		path := []complex64{start}
		next := start + direction
		for next != next+direction && next != start {
			path = append(path, next)
			next, direction = next+direction, pipeDirection(direction, grid[next+direction])
		}
		return path
	}
	loop := []complex64{}
	start := grid.FindAny('S')
	for _, dir := range grid.Directions() {
		if lp := findLoop(start, dir); len(lp) > len(loop) {
			loop = lp
		}
	}
	solution.Part1 = (len(loop) + 1) / 2
	for c, row := range input {
		in := false
		previousturn := ' '
		for r, v := range row {
			if slices.Contains(loop, complex(float32(r), float32(c))) {
				if v == '|' ||
					v == '7' && previousturn == 'L' ||
					v == 'J' && previousturn == 'F' {
					in = !in
				}
				if strings.ContainsRune("7LJF", v) {
					previousturn = v
				}
			} else if in {
				solution.Part2++
			}
		}
	}
	return
}

func pipeDirection(directionfrom complex64, r rune) complex64 {
	switch directionfrom {
	case -1i:
		switch r {
		case '|':
			return -1i
		case '7':
			return -1
		case 'F':
			return 1
		}
	case 1i:
		switch r {
		case '|':
			return 1i
		case 'J':
			return -1
		case 'L':
			return 1
		}
	case 1:
		switch r {
		case '-':
			return 1
		case '7':
			return 1i
		case 'J':
			return -1i
		}
	case -1:
		switch r {
		case '-':
			return -1
		case 'L':
			return -1i
		case 'F':
			return 1i
		}
	}
	return 0
}
