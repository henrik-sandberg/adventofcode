package y2025

import (
	"adventofcode/solutions/shared"
)

func Day07(input []string) (solution shared.Solution[int, int]) {
	grid := shared.NewGrid(input)
	memo := make(map[complex128]int)
	var dfs func(c complex128) int
	dfs = func(c complex128) int {
		if val, ok := memo[c]; ok {
			return val
		}
		if _, ok := grid[c]; !ok {
			return 1
		}
		left := c - 1
		for ; grid[left] != 0 && grid[left] != '^'; left += 1i {
		}
		right := c + 1
		for ; grid[right] != 0 && grid[right] != '^'; right += 1i {
		}
		total := dfs(left) + dfs(right)
		memo[c] = total
		return total
	}
	solution.Part2 = dfs(grid.FindAny('S'))
	solution.Part1 = len(memo)
	return
}
