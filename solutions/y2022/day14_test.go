package y2022

import (
	"strings"
	"testing"
)

func TestDay14(t *testing.T) {
	sampleInput := `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`

	input := strings.Split(sampleInput, "\n")

	res := Day14(input)

	if expected := 24; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 93; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
