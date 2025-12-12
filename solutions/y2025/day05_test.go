package y2025

import (
	"strings"
	"testing"
)

func TestDay05(t *testing.T) {
	sampleInput := `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

	input := strings.Split(sampleInput, "\n")

	res := Day05(input)

	if expected := 3; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 14; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
