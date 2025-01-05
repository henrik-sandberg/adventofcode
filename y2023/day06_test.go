package y2023

import (
	"strings"
	"testing"
)

func TestDay06(t *testing.T) {
	sampleInput := strings.Split(strings.TrimSpace(`
Time:      7  15   30
Distance:  9  40  200
`), "\n")

	res := Day06(sampleInput)

	if expected := 288; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 71503; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
