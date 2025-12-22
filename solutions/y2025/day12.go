package y2025

import (
	"adventofcode/solutions/shared"
	"strings"
)

func Day12(input []string) (solution shared.Solution[int, int]) {
	var pixelCount []int
	idx := 0
	for ; !strings.ContainsRune(input[idx], 'x'); idx += 5 {
		s := strings.Join(input[idx+1:idx+4], "")
		pixelCount = append(pixelCount, strings.Count(s, "#"))
	}
	for ; idx < len(input); idx++ {
		fs := strings.Fields(input[idx])
		capacity := shared.Product(shared.IntSlice(strings.Split(strings.Trim(fs[0], ":"), "x"))...)
		pixels := 0
		for i, v := range shared.IntSlice(fs[1:]) {
			pixels += pixelCount[i] * v
		}
		if pixels <= capacity {
			solution.Part1 += 1
		}
	}
	return
}
