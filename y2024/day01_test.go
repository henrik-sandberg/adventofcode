package y2024

import (
	"strings"
	"testing"
)

func TestDay01(t *testing.T) {
	sampleInput := `3   4
4   3
2   5
1   3
3   9
3   3`

	input := strings.Split(sampleInput, "\n")

	res := Day01(input)

	if expected := 11; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 31; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
