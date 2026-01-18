package y2022

import (
	"adventofcode/solutions/shared"
	"math"
	"slices"
)

func Day23(input []string) (solution shared.Solution[int, int]) {
	grid := make(map[complex128]bool)
	for k, v := range shared.NewGrid(input) {
		if v == '#' {
			grid[k] = true
		}
	}
	adj := []complex128{-1 - 1i, -1i, 1 - 1i, 1, 1 + 1i, 1i, -1 + 1i, -1}
	front := [][]complex128{
		{-1i, -1 - 1i, 1 - 1i},
		{+1i, -1 + 1i, 1 + 1i},
		{-1, -1 - 1i, -1 + 1i},
		{+1, +1 - 1i, +1 + 1i},
	}
	findMove := func(p complex128, round int) complex128 {
		for i := range 4 {
			dir := front[(round+i)%4]
			if !slices.ContainsFunc(dir, func(d complex128) bool {
				return grid[p+d]
			}) {
				return p + dir[0]
			}
		}
		return p
	}
	for round := 0; ; round++ {
		stale := 0
		candidates := make(map[complex128][]complex128)
		for p := range grid {
			if !slices.ContainsFunc(adj, func(d complex128) bool {
				return grid[p+d]
			}) {
				stale++
				continue
			}
			c := findMove(p, round)
			candidates[c] = append(candidates[c], p)
		}
		if stale == len(grid) {
			solution.Part2 = round + 1
			break
		}
		for to, from := range candidates {
			if len(from) == 1 {
				delete(grid, from[0])
				grid[to] = true
			}
		}
		if round == 10 {
			up := math.MaxInt
			down := math.MinInt
			left := math.MaxInt
			right := math.MinInt
			for k := range grid {
				up = min(up, int(imag(k)))
				down = max(down, int(imag(k)))
				left = min(left, int(real(k)))
				right = max(right, int(real(k)))
			}
			solution.Part1 = (down-up+1)*(right-left+1) - len(grid)
		}
	}
	return
}
