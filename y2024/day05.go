package y2024

import (
	"adventofcode/shared"
	"slices"
	"strconv"
	"strings"
)

func Day05(input []string) shared.Solution {
	orderIndex := slices.Index(input, "")
	rules := map[string][]string{}
	for _, rule := range input[:orderIndex] {
		r := strings.Split(rule, "|")
		rules[r[1]] = append(rules[r[1]], r[0])
	}
	orderSortFunc := func(a, b string) int {
		if slices.Contains(rules[a], b) {
			return 1
		}
		if slices.Contains(rules[b], a) {
			return -1
		}
		return 0
	}
	part1 := 0
	part2 := 0
	for _, order := range input[orderIndex+1:] {
		ord := strings.Split(order, ",")
		sorted := make([]string, len(ord))
		copy(sorted, ord)
		slices.SortStableFunc(sorted, orderSortFunc)
		pageValue, _ := strconv.Atoi(sorted[len(sorted)/2])
		if slices.Equal(ord, sorted) {
			part1 += pageValue
		} else {
			part2 += pageValue
		}
	}
	return shared.Solution{Part1: part1, Part2: part2}
}
