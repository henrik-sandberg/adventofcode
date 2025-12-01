package y2021

import (
	"strings"
	"testing"
)

func TestDay03(t *testing.T) {
	sampleInput := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

	input := strings.Split(sampleInput, "\n")

	res := Day03(input)

	if expected := 198; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 230; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
