package main

import (
	"fmt"
	"strings"
)

func Day10(input []string) {
	grid := Map{strings.Join(input, ""), len(input[0])}.ToComplexGrid()
	solve := func(part1 bool) (res int) {
		for start, val := range grid {
			if val == '0' {
				seen := map[complex64]bool{}
				queue := []complex64{start}
				for len(queue) > 0 {
					p := queue[0]
					queue = queue[1:]
					for _, dir := range []complex64{1, -1, -1i, 1i} {
						if next := p + dir; grid[p]+1 == grid[next] && !seen[next] {
							seen[next] = part1
							if grid[next] == '9' {
								res++
							} else {
								queue = append(queue, next)
							}
						}
					}
				}
			}
		}
		return
	}
	fmt.Println(solve(true))
	fmt.Println(solve(false))
}
