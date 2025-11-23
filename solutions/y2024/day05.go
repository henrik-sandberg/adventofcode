package y2024

import (
	"adventofcode/solutions/shared"
	"slices"
	"strconv"
	"strings"
)

func Day05(input []string) (solution shared.Solution[int, int]) {
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
	for _, order := range input[orderIndex+1:] {
		ord := strings.Split(order, ",")
		sorted := make([]string, len(ord))
		copy(sorted, ord)
		slices.SortStableFunc(sorted, orderSortFunc)
		pageValue, _ := strconv.Atoi(sorted[len(sorted)/2])
		if slices.Equal(ord, sorted) {
			solution.Part1 += pageValue
		} else {
			solution.Part2 += pageValue
		}
	}
	return
}
