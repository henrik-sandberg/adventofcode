package y2021

import (
	"strings"
	"testing"
)

func TestDay02(t *testing.T) {
	sampleInput := `forward 5
down 5
forward 8
up 3
down 8
forward 2`

	input := strings.Split(sampleInput, "\n")

	res := Day02(input)

	if expected := 150; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 900; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
