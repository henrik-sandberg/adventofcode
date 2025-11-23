package y2024

import (
	"strings"
	"testing"
)

func TestDay02(t *testing.T) {
	sampleInput := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	input := strings.Split(sampleInput, "\n")

	res := Day02(input)

	if expected := 2; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 4; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
