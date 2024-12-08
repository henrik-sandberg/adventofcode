package main

import (
	"fmt"
	"slices"
	"strings"
)

func Day05(input []string) {
	ind := IndexOf(input, "")
	rules := map[string][]string{}
	for _, rule := range input[:ind] {
		r := strings.Split(rule, "|")
		rules[r[1]] = append(rules[r[1]], r[0])
	}
	part1 := 0
	part2 := 0
	for _, order := range input[ind+1:] {
		ord := strings.Split(order, ",")
		sorted := ordersSorted(ord, rules)
		pageValue := ToInts(sorted)[len(sorted)/2]
		if slices.Equal(ord, sorted) {
			part1 += pageValue
		} else {
			part2 += pageValue
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

func ordersSorted(order []string, rules map[string][]string) []string {
	s := make([]string, len(order))
	copy(s, order)
	slices.SortStableFunc(s, func(a, b string) int {
		if slices.Contains(rules[a], b) {
			return 1
		}
		if slices.Contains(rules[b], a) {
			return -1
		}
		return 0
	})
	return s
}
