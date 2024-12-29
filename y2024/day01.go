package y2024

import (
	"adventofcode/shared"
	"sort"
	"strings"
)

func Day01(input []string) shared.Solution {
	parsed := [][]int{{}, {}}
	for _, line := range input {
		ints := shared.IntSlice(strings.Fields(line))
		parsed[0] = append(parsed[0], ints[0])
		parsed[1] = append(parsed[1], ints[1])
	}
	sort.Sort(sort.IntSlice(parsed[0]))
	sort.Sort(sort.IntSlice(parsed[1]))
	part1 := 0
	for i := range input {
		part1 = part1 + shared.Abs(parsed[0][i]-parsed[1][i])
	}
	part2 := 0
	for _, v := range parsed[0] {
		part2 = part2 + v*shared.Count(parsed[1], v)
	}
	return shared.Solution{Part1: part1, Part2: part2}
}
