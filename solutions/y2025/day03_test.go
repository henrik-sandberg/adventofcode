package y2025

import (
	"strings"
	"testing"
)

func TestDay03(t *testing.T) {
	sampleInput := `987654321111111
811111111111119
234234234234278
818181911112111`

	input := strings.Split(sampleInput, "\n")

	res := Day03(input)

	if expected := 357; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 3121910778619; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
