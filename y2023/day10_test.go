package y2023

import (
	"strings"
	"testing"
)

func TestDay10(t *testing.T) {
	sampleInput := `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`

	res := Day10(strings.Split(sampleInput, "\n"))

	if expected := 8; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 0; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
