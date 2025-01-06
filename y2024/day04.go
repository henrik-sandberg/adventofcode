package y2024

import (
	"adventofcode/shared"
	"strings"
)

func Day04(input []string) (solution shared.Solution[int, int]) {
	h, w := len(input), len(input[0])
	rows := make([]strings.Builder, h)
	cols := make([]strings.Builder, w)
	fdiag := make([]strings.Builder, h+w-1)
	bdiag := make([]strings.Builder, len(fdiag))
	for y, row := range input {
		for x, val := range row {
			rows[y].WriteRune(val)
			cols[x].WriteRune(val)
			fdiag[x+y].WriteRune(val)
			bdiag[x-y+h-1].WriteRune(val)
		}
	}
	search := "XMAS"
	searchReverse := shared.Reverse(search)
	for _, arr := range [][]strings.Builder{rows, cols, fdiag, bdiag} {
		for _, sb := range arr {
			solution.Part1 += strings.Count(sb.String(), search) + strings.Count(sb.String(), searchReverse)
		}
	}
	grid := shared.NewGrid(input)
	for _, a := range grid.FindAll('A') {
		ns := []byte{
			grid[a-1-1i],
			grid[a+1-1i],
			grid[a-1+1i],
			grid[a+1+1i],
		}
		solution.Part2 += shared.BoolToInt(ns[0] != ns[3] && shared.Count(ns, 'M') == 2 && shared.Count(ns, 'S') == 2)
	}
	return
}
