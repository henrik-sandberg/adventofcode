package y2025

import (
	"strings"
	"testing"
)

func TestDay09(t *testing.T) {
	sampleInput := `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`

	input := strings.Split(sampleInput, "\n")

	res := Day09(input)

	if expected := 50; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 24; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
