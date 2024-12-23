package main

import (
	"fmt"
	"slices"
	"strings"
)

func Day23(input []string) {
	edges := map[string][]string{}
	for _, edge := range input {
		e := strings.Split(edge, "-")
		edges[e[0]] = append(edges[e[0]], e[1])
		edges[e[1]] = append(edges[e[1]], e[0])
	}
	isRing := func(s []string) bool {
		for pair := range Combinations(s, 2) {
			if !slices.Contains(edges[pair[0]], pair[1]) {
				return false
			}
		}
		return true
	}
	rings := map[string]bool{}
	for k, v := range edges {
		if strings.HasPrefix(k, "t") {
			for combination := range Combinations(v, 2) {
				if isRing(combination) {
					ring := append(combination, k)
					slices.Sort(ring)
					rings[strings.Join(ring, ",")] = true
				}
			}
		}
	}
	fmt.Println(len(rings))
	maxLength := 0
	for _, v := range edges {
		if len(v) > maxLength {
			maxLength = len(v)
		}
	}
	fmt.Println(func() string {
		for length := maxLength; ; length-- {
			for k, v := range edges {
				for combination := range Combinations(v, length) {
					if isRing(combination) {
						ring := append(combination, k)
						slices.Sort(ring)
						return strings.Join(ring, ",")
					}

				}
			}
		}
	}())
}
