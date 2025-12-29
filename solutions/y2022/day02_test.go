package y2022

import (
	"strings"
	"testing"
)

func TestDay02(t *testing.T) {
	sampleInput := `A Y
B X
C Z`

	input := strings.Split(sampleInput, "\n")

	res := Day02(input)

	if expected := 15; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 12; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
