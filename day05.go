package main

import (
	"fmt"
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
		if isValidOrder(ord, rules) {
			part1 += ToInts(ord)[len(ord)/2]
		} else {
			fixed := fixInvalidOrder(ord, rules)
			part2 += ToInts(fixed)[len(ord)/2]
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

func isValidOrder(order []string, rules map[string][]string) bool {
	printed := map[string]bool{}
	for _, page := range order {
		for _, needed := range rules[page] {
			if !printed[needed] && Contains(order, needed) {
				return false
			}
		}
		printed[page] = true
	}
	return true
}

func fixInvalidOrder(order []string, rules map[string][]string) []string {
	ret := []string{}
	for len(ret) < len(order) {
	findPage:
		for _, page := range order {
			if !Contains(ret, page) {
				for _, needed := range rules[page] {
					if !Contains(ret, needed) && Contains(order, needed) {
						continue findPage
					}
				}
				ret = append(ret, page)
			}
		}
	}
	return ret
}
