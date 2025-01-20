package y2023

import (
	"adventofcode/shared"
	"strings"
)

func Day11(input []string) (solution shared.Solution[int, int]) {
	solver := func(expansionMultiplier int) int {
		multiplier := max(1, expansionMultiplier-1)
		rowadjust := []int{}
		for i, row := range input {
			previous := 0
			if len(rowadjust) > 0 {
				previous += rowadjust[i-1]
			}
			rowadjust = append(rowadjust, previous+multiplier*shared.BoolToInt(!strings.ContainsRune(row, '#')))
		}
		coladjust := []int{}
		for j := range input[0] {
			previous := 0
			if len(coladjust) > 0 {
				previous += coladjust[j-1]
			}
			hasGalaxy := false
			for i := range input {
				if input[i][j] == '#' {
					hasGalaxy = true
					break
				}
			}
			coladjust = append(coladjust, previous+multiplier*shared.BoolToInt(!hasGalaxy))
		}
		grid := shared.Grid{}
		for ri, row := range input {
			for ci, cell := range []byte(row) {
				if cell == '#' {
					grid[complex(float64(ci+coladjust[ci]), float64(ri+rowadjust[ri]))] = cell
				}
			}
		}
		res := 0
		for pair := range shared.Combinations(grid.FindAll('#'), 2) {
			a, b := pair[0], pair[1]
			res += shared.Abs(int(real(a)-real(b))) + shared.Abs(int(imag(a)-imag(b)))
		}
		return res
	}
	solution.Part1 = solver(1)
	solution.Part2 = solver(1000000)
	return
}
