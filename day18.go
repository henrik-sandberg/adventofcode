package main

import (
	"fmt"
	"strings"
)

func Day18(input []string) {
	iterations := 1024
	size := 71
	target := complex64(70 + 70i)
	m := map[complex64]rune{}
	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			m[complex(float32(c), float32(r))] = '.'
		}
	}
	dropByte := func(point string) {
		i := ToInts(strings.Split(point, ","))
		m[complex(float32(i[0]), float32(i[1]))] = '#'
	}
	bfs := func() int {
		current := complex64(0)
		queue := []complex64{current}
		seen := map[complex64]int{current: 0}
		for len(queue) > 0 {
			current, queue = queue[0], queue[1:]
			for _, dir := range []complex64{-1i, 1, 1i, -1} {
				next := current + dir
				if next == target {
					return seen[current] + 1
				}
				_, visited := seen[next]
				if !visited && m[next] == '.' {
					queue = append(queue, next)
					seen[next] = seen[current] + 1
				}
			}
		}
		return -1
	}
	for _, line := range input[:iterations] {
		dropByte(line)
	}
	fmt.Println(bfs())
	for _, line := range input[iterations:] {
		dropByte(line)
		if bfs() == -1 {
			fmt.Println(line)
			break
		}
	}
}
