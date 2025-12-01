package y2025

import (
	"strings"
	"testing"
)

func TestDay01(t *testing.T) {
	sampleInput := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

	input := strings.Split(sampleInput, "\n")

	res := Day01(input)

	if expected := 3; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 6; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
