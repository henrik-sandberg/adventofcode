package y2022

import (
	"adventofcode/solutions/shared"
	"maps"
	"regexp"
	"slices"
	"strconv"
)

func Day22(input []string) (solution shared.Solution[int, int]) {
	splitIdx := slices.Index(input, "")
	moves := regexp.MustCompile(`\d+|\w`).FindAllString(input[splitIdx+1], -1)
	grid := shared.NewGrid(input[:splitIdx])
	maps.DeleteFunc(grid, func(_ complex128, v byte) bool {
		return v == ' '
	})

	solve := func(stepFunc func(complex128, complex128) (complex128, complex128)) int {
		pos := 0i
		dir := 1 + 0i
		for {
			_, found := grid[pos]
			if found {
				break
			}
			pos += dir
		}
		for _, move := range moves {
			switch move {
			case "R":
				dir *= 1i
			case "L":
				dir *= -1i
			default:
				step, _ := strconv.Atoi(move)
				for range step {
					pos, dir = stepFunc(pos, dir)
				}
			}
		}
		var dirValue int
		switch dir {
		case 1i:
			dirValue = 1
		case -1:
			dirValue = 2
		case -1i:
			dirValue = 3
		}
		return int(1000*(imag(pos)+1)+4*(real(pos)+1)) + dirValue
	}
	solution.Part1 = solve(func(start, dir complex128) (complex128, complex128) {
		pos := start + dir
		_, found := grid[pos]
		if !found {
			for {
				_, found := grid[pos-dir]
				if !found {
					break
				}
				pos -= dir
			}
		}
		if grid[pos] == '#' {
			return start, dir
		}
		return pos, dir
	})
	return
}
