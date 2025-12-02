package y2021

import (
	"strings"
	"testing"
)

func TestDay05(t *testing.T) {
	sampleInput := `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

	input := strings.Split(sampleInput, "\n")

	res := Day05(input)

	if expected := 5; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 12; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
