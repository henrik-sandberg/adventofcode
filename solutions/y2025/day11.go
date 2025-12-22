package y2025

import (
	"adventofcode/solutions/shared"
	"strings"
)

func Day11(input []string) (solution shared.Solution[int, int]) {
	graph := make(map[string][]string)
	for _, line := range input {
		fs := strings.Fields(line)
		graph[strings.Trim(fs[0], ":")] = fs[1:]
	}
	var dfs func(string, map[string]int) int
	dfs = func(node string, memo map[string]int) int {
		if v, ok := memo[node]; ok {
			return v
		}
		res := 0
		for _, next := range graph[node] {
			res += dfs(next, memo)
		}
		memo[node] = res
		return res
	}
	countPaths := func(nodes ...string) int {
		res := 1
		for idx := range len(nodes) - 1 {
			res *= dfs(nodes[idx], map[string]int{nodes[idx+1]: 1})
		}
		return res
	}
	solution.Part1 = countPaths("you", "out")
	solution.Part2 = countPaths("svr", "fft", "dac", "out") + countPaths("svr", "dac", "fft", "out")
	return
}
