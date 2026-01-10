package y2022

import (
	"strings"
	"testing"
)

func TestDay18(t *testing.T) {
	sampleInput := `2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5`

	input := strings.Split(sampleInput, "\n")

	res := Day18(input)

	if expected := 64; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 58; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
