package y2025

import (
	"strings"
	"testing"
)

func TestDay01(t *testing.T) {
	sampleInput := ""

	input := strings.Split(sampleInput, "\n")

	res := Day01(input)

	if expected := 0; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 0; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
