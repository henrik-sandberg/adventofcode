package y2025

import (
	"adventofcode/solutions/shared"
	"slices"
	"strconv"
	"strings"
)

func Day05(input []string) (solution shared.Solution[int, int]) {
	type Range struct {
		from, to int
	}
	var ranges []Range
	idx := 0
	for ; input[idx] != ""; idx++ {
		s := strings.Split(input[idx], "-")
		f, _ := strconv.Atoi(s[0])
		t, _ := strconv.Atoi(s[1])
		ranges = append(ranges, Range{from: f, to: t})
	}
	slices.SortFunc(ranges, func(a, b Range) int {
		return a.from - b.from
	})
	mergedRanges := []Range{ranges[0]}
	for _, r := range ranges {
		last := mergedRanges[len(mergedRanges)-1]
		if r.from > last.to {
			mergedRanges = append(mergedRanges, r)
		} else {
			mergedRanges[len(mergedRanges)-1] = Range{from: last.from, to: max(last.to, r.to)}
		}
	}
	for _, v := range shared.IntSlice(input[idx+1:]) {
		if slices.ContainsFunc(mergedRanges, func(r Range) bool {
			return r.from <= v && v <= r.to

		}) {
			solution.Part1++
		}
	}
	for _, r := range mergedRanges {
		solution.Part2 += r.to - r.from + 1
	}
	return
}
