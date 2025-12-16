package y2024

import (
	"adventofcode/solutions/shared"
)

func Day20(input []string) shared.Solution[int, int] {
	return day20Solver(input, 100)
}

func day20Solver(input []string, limit int) (solution shared.Solution[int, int]) {
	grid := shared.NewGrid(input)
	start := grid.FindAny('S')
	target := grid.FindAny('E')
	arr := []complex128{start}
	for arr[len(arr)-1] != target {
		for _, dir := range []complex128{-1i, 1, 1i, -1} {
			next := arr[len(arr)-1] + dir
			if grid[next] != '#' && (len(arr) < 2 || arr[len(arr)-2] != next) {
				arr = append(arr, next)
				break
			}
		}
	}
	manhattan := func(a, b complex128) int {
		return shared.Abs(int(real(a)-real(b))) + shared.Abs(int(imag(a)-imag(b)))
	}
	for i, first := range arr {
		for j := i + limit; j < len(arr); j++ {
			cheatDistance := manhattan(first, arr[j])
			saved := (j - i) - cheatDistance
			if cheatDistance <= 2 && saved >= limit {
				solution.Part1++
			}
			if cheatDistance <= 20 && saved >= limit {
				solution.Part2++
			}
		}
	}
	return
}
