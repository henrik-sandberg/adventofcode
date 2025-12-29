package y2022

import (
	"strings"
	"testing"
)

func TestDay04(t *testing.T) {
	sampleInput := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

	input := strings.Split(sampleInput, "\n")

	res := Day04(input)

	if expected := 2; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 4; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
