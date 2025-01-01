package y2022

import (
	"adventofcode/shared"
	"sort"
	"strconv"
)

func Day01(input []string) (solution shared.Solution[int, int]) {
	counters := make([]int, 1)
	for _, s := range input {
		if s == "" {
			counters = append(counters, 0)
		} else {
			i, _ := strconv.Atoi(s)
			counters[len(counters)-1] += i
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counters)))
	solution.Part1 = counters[0]
	solution.Part2 = counters[0] + counters[1] + counters[2]
	return
}
