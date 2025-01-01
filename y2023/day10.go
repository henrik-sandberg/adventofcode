package y2023

import (
	"adventofcode/shared"
)

func Day10(input []string) (solution shared.Solution[int, int]) {
	grid := shared.NewGrid(input)
	findLoop := func(start complex64, direction complex64) []complex64 {
		path := []complex64{}
		next := start + direction
		direction = pipeDirection(direction, grid[next])
		for next != next+direction && next != start {
			next, direction = next+direction, pipeDirection(direction, grid[next+direction])
			path = append(path, next)
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
