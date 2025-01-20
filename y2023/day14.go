package y2023

import (
	"adventofcode/shared"
	"fmt"
	"slices"
)

func Day14(input []string) (solution shared.Solution[int, int]) {
	grid := shared.NewGrid(input)
	roll := func(dir complex128) int {
		rocks := grid.FindAll('O')
		slices.SortFunc(rocks, func(a, b complex128) int {
			switch {
			case real(dir) > 0:
				return int(real(b) - real(a))
			case real(dir) < 0:
				return int(real(a) - real(b))
			case imag(dir) > 0:
				return int(imag(b) - imag(a))
			case imag(dir) < 0:
				return int(imag(a) - imag(b))
			default:
				return 0
			}
		})
		swaps := 0
		for _, n := range rocks {
			next := n + dir
			if grid[next] == '.' {
				swaps++
				grid[n], grid[next] = grid[next], grid[n]
			}
		}
		return swaps
	}
	score := func() (res int) {
		for k, v := range grid {
			if v == 'O' {
				res += len(input) - int(imag(k))
			}
		}
		return res
	}
	for roll(grid.Up()) != 0 {
	}
	solution.Part1 = score()
	seen := map[string]int{}
	scores := []int{}
	for cycle := 0; ; cycle++ {
		for _, dir := range []complex128{grid.Up(), grid.Left(), grid.Down(), grid.Right()} {
			for roll(dir) != 0 {
			}
		}
		key := fmt.Sprint(grid)
		if previous, ok := seen[key]; ok {
			offset := (1e9 - previous) % (cycle - previous)
			solution.Part2 = scores[previous+offset-1]
			return
		}
		seen[key] = cycle
		scores = append(scores, score())
	}
}
