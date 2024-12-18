package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

/*
	Program: 2,4,1,5,7,5,1,6,0,3,4,3,5,5,3,0

	while a != 0:
		b = a % 8
		b = b ^ 5
		c = a / (1 << b)
		b = b ^ 6
		a = a / 8
		b = b ^ c
		out b % 8
*/

func Day17(input []string) {
	registerA, _ := strconv.Atoi(strings.Split(input[0], " ")[2])
	program := ToInts(strings.Split(strings.Split(input[4], " ")[1], ","))
	slices.Reverse(program)
	out := func(a int) int {
		return a&7 ^ 3 ^ a/(1<<(a&7^5))&7
	}
	output := []string{}
	for registerA != 0 {
		output = append(output, fmt.Sprint(out(registerA)))
		registerA >>= 3
	}
	fmt.Println(strings.Join(output, ","))

	var dfs func(seen []int, a int) int
	dfs = func(seen []int, a int) int {
		if slices.Equal(seen, program) {
			return a
		}
		wanted := program[len(seen)]
		for n := a * 8; n < (a+1)*8; n++ {
			if out(n) == wanted {
				seen = append(seen, wanted)
				if ans := dfs(seen, n); ans > 0 {
					return ans
				}
				seen = seen[:len(seen)-1]
			}
		}
		return 0
	}
	fmt.Println(dfs([]int{}, 0))
}
