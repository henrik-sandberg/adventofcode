package main

import (
	"fmt"
	"strconv"
	"strings"
)

type stoneSplitLength struct {
	stone     string
	remaining int
}

func Day11(input []string) {
	part1, part2 := 0, 0
	memo := map[stoneSplitLength]int{}
	for _, stone := range strings.Fields(input[0]) {
		part1 += split(memo, stone, 0, 25)
		part2 += split(memo, stone, 0, 75)
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

func split(memo map[stoneSplitLength]int, v string, current, depth int) (res int) {
	key := stoneSplitLength{v, depth - current}
	if v, ok := memo[key]; ok {
		return v
	}
	if current == depth {
		return 1
	}
	switch {
	case v == "0":
		res = split(memo, "1", current+1, depth)
	case len(v)%2 == 0:
		left, _ := strconv.Atoi(v[:len(v)/2])
		right, _ := strconv.Atoi(v[len(v)/2:])
		res = split(memo, fmt.Sprint(left), current+1, depth) + split(memo, fmt.Sprint(right), current+1, depth)
	default:
		val, _ := strconv.Atoi(v)
		res = split(memo, fmt.Sprint(val*2024), current+1, depth)
	}
	memo[key] = res
	return
}
