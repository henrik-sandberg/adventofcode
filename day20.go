package main

import (
	"fmt"
	"strings"
)

func Day20(input []string) {
	m := Map{strings.Join(input, ""), len(input[0])}.ToComplexGrid()
	var start, target complex64
	for k, v := range m {
		switch v {
		case 'S':
			start = k
		case 'E':
			target = k
		}
	}
	arr := []complex64{start}
	for arr[len(arr)-1] != target {
		for _, dir := range []complex64{-1i, 1, 1i, -1} {
			next := arr[len(arr)-1] + dir
			if m[next] != '#' && (len(arr) < 2 || arr[len(arr)-2] != next) {
				arr = append(arr, next)
				break
			}
		}
	}
	part1 := 0
	part2 := 0
	limit := 100
	manhattan := func(a, b complex64) int {
		return abs(int(real(a)-real(b))) + abs(int(imag(a)-imag(b)))
	}
	for i, first := range arr {
		for j := i + limit; j < len(arr); j++ {
			cheatDistance := manhattan(first, arr[j])
			saved := (j - i) - cheatDistance
			if cheatDistance <= 2 && saved >= limit {
				part1++
			}
			if cheatDistance <= 20 && saved >= limit {
				part2++
			}
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
}
