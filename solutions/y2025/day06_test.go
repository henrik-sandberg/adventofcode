package y2025

import (
	"strings"
	"testing"
)

func TestDay06(t *testing.T) {
	sampleInput := `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `

	input := strings.Split(sampleInput, "\n")

	res := Day06(input)

	if expected := 4277556; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 3263827; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
