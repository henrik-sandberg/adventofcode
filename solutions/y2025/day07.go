package y2025

import (
	"adventofcode/solutions/shared"
)

func Day07(input []string) (solution shared.Solution[int, int]) {
	grid := shared.NewGrid(input)
	findNextNode := func(p complex128) complex128 {
		for grid[p] != 0 {
			if grid[p] == '^' {
				return p
			}
			p += 1i
		}
		return -1
	}
	var memo = make(map[complex128]int)
	var dfs func(c complex128) int
	dfs = func(c complex128) int {
		if val, ok := memo[c]; ok {
			return val
		}
		if _, ok := grid[c]; !ok {
			return 1
		}
		left, right := findNextNode(c-1), findNextNode(c+1)
		total := dfs(left) + dfs(right)
		memo[c] = total
		return total
	}
	start := grid.FindAny('S')
	dfs(start)
	solution.Part1 = len(memo)
	solution.Part2 = memo[start]
	return
}
