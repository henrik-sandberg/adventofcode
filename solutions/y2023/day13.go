package y2023

import (
	"adventofcode/solutions/shared"
	"slices"
)

func Day13(input []string) (solution shared.Solution[int, int]) {
	score := func(miss, val int) {
		switch miss {
		case 0:
			solution.Part1 += val
		case 1:
			solution.Part2 += val
		}
	}
	vertical := func(pattern []string) {
		h, w := len(pattern), len(pattern[0])
		for x := 0; x < w; x++ {
			miss := 0
			for offset := 1; x-offset >= 0 && x+offset-1 < w; offset++ {
				for ri := 0; ri < h; ri++ {
					miss += shared.BoolToInt(pattern[ri][x-offset] != pattern[ri][x+offset-1])
				}
			}
			score(miss, x)
		}
	}
	horizontal := func(pattern []string) {
		h, w := len(pattern), len(pattern[0])
		for y := 0; y < h; y++ {
			miss := 0
			for offset := 1; y-offset >= 0 && y+offset-1 < h; offset++ {
				for ci := 0; ci < w; ci++ {
					miss += shared.BoolToInt(pattern[y-offset][ci] != pattern[y+offset-1][ci])
				}
			}
			score(miss, 100*y)
		}
	}
	for len(input) > 0 {
		ind := slices.Index(input, "")
		if ind == -1 {
			ind = len(input) - 1
		}
		vertical(input[:ind])
		horizontal(input[:ind])
		input = input[ind+1:]
	}
	return
}
