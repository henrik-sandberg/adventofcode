package y2022

import (
	"strings"
	"testing"
)

func TestDay09(t *testing.T) {
	sampleInput := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

	input := strings.Split(sampleInput, "\n")

	res := Day09(input)

	if expected := 13; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 1; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
