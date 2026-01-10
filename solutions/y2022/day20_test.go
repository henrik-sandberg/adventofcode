package y2022

import (
	"strings"
	"testing"
)

func TestDay20(t *testing.T) {
	sampleInput := `1
2
-3
3
-2
0
4`

	input := strings.Split(sampleInput, "\n")

	res := Day20(input)

	if expected := 3; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 1623178306; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
