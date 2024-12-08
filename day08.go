package main

import (
	"fmt"
	"strings"
)

func Day08(input []string) {
	grid := Map{strings.Join(input, ""), len(input[0])}.ToComplexGrid()
	frequencies := Bag(strings.Join(input, ""))
	delete(frequencies, '.')
	for _, limit := range []struct{ from, to int }{{1, 2}, {0, 50}} {
		antinodes := map[complex64]bool{}
		for freq := range frequencies {
			towers := []complex64{}
			for k, v := range grid {
				if v == freq {
					towers = append(towers, k)
				}
			}
			for nums := range Permutations(towers, 2) {
				a, b := nums[0], nums[1]
				for i := limit.from; i < limit.to; i++ {
					n := complex(float32(i), 0)
					antinodes[a+n*(a-b)] = true
				}
			}
		}
		ans := 0
		for a := range antinodes {
			if _, ok := grid[a]; ok {
				ans++
			}
		}
		fmt.Println(ans)
	}
}
