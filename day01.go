package main

import (
	"fmt"
	"sort"
	"strings"
)

func Day01(input []string) {
	parsed := [][]int{{}, {}}
	for _, line := range input {
		ints := ToInts(strings.Fields(line))
		parsed[0] = append(parsed[0], ints[0])
		parsed[1] = append(parsed[1], ints[1])
	}
	sort.Sort(sort.IntSlice(parsed[0]))
	sort.Sort(sort.IntSlice(parsed[1]))
	part1 := 0
	for i := range input {
		part1 = part1 + abs(parsed[0][i]-parsed[1][i])
	}
	fmt.Println(part1)

	part2 := 0
	for _, v := range parsed[0] {
		part2 = part2 + v*Count(parsed[1], v)
	}
	fmt.Println(part2)
}
