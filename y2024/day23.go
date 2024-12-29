package y2024

import (
	"adventofcode/shared"
	"slices"
	"strings"
)

func Day23(input []string) shared.Solution {
	edges := map[string][]string{}
	for _, edge := range input {
		e := strings.Split(edge, "-")
		edges[e[0]] = append(edges[e[0]], e[1])
		edges[e[1]] = append(edges[e[1]], e[0])
	}
	allConnected := func(s []string) bool {
		for pair := range shared.Combinations(s, 2) {
			if !slices.Contains(edges[pair[0]], pair[1]) {
				return false
			}
		}
		return true
	}
	rings := map[string]bool{}
	for k, v := range edges {
		if strings.HasPrefix(k, "t") {
			for combination := range shared.Combinations(v, 2) {
				if allConnected(combination) {
					ring := append(combination, k)
					slices.Sort(ring)
					rings[strings.Join(ring, ",")] = true
				}
			}
		}
	}
	part1 := len(rings)
	part2 := func() string {
		length := 0
		for _, v := range edges {
			if len(v) > length {
				length = len(v)
			}
		}
		for ; ; length-- {
			for k, v := range edges {
				for combination := range shared.Combinations(v, length) {
					if allConnected(combination) {
						ans := append(combination, k)
						slices.Sort(ans)
						return strings.Join(ans, ",")
					}
				}
			}
		}
	}()
	return shared.Solution{Part1: part1, Part2: part2}
}
