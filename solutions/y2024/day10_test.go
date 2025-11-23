package y2024

import (
	"strings"
	"testing"
)

func TestDay10(t *testing.T) {
	sampleInput := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

	input := strings.Split(sampleInput, "\n")

	res := Day10(input)

	if expected := 36; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 81; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
