package y2023

import (
	"adventofcode/solutions/shared"
	"slices"
	"strings"
)

func Day10(input []string) (solution shared.Solution[int, int]) {
	grid := shared.NewGrid(input)
	findLoop := func(start complex128, direction complex128) []complex128 {
		path := []complex128{start}
		next := start + direction
		for next != next+direction && next != start {
			path = append(path, next)
			next, direction = next+direction, pipeDirection(direction, grid[next+direction])
		}
		return path
	}
	loop := []complex128{}
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
			if slices.Contains(loop, complex(float64(r), float64(c))) {
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

func pipeDirection(directionfrom complex128, b byte) complex128 {
	switch directionfrom {
	case -1i:
		switch b {
		case '|':
			return -1i
		case '7':
			return -1
		case 'F':
			return 1
		}
	case 1i:
		switch b {
		case '|':
			return 1i
		case 'J':
			return -1
		case 'L':
			return 1
		}
	case 1:
		switch b {
		case '-':
			return 1
		case '7':
			return 1i
		case 'J':
			return -1i
		}
	case -1:
		switch b {
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
