package main

import (
	"fmt"
	"strings"
)

func Day19(input []string) {
	available := strings.Split(input[0], ", ")
	cache := map[string]int{}
	var search func(string, []string) int
	search = func(s string, work []string) (res int) {
		if len(s) == 0 {
			return 1
		}
		if v, ok := cache[s]; ok {
			return v
		}
		for _, a := range available {
			if strings.HasPrefix(s, a) {
				work = append(work, a)
				res += search(s[len(a):], work)
				work = work[:len(work)-1]
			}
		}
		cache[s] = res
		return res
	}
	part1 := 0
	part2 := 0
	for _, design := range input[2:] {
		found := search(design, []string{})
		if found > 0 {
			part1 += 1
		}
		part2 += found

	}
	fmt.Println(part1)
	fmt.Println(part2)
}
