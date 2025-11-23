package y2021

import (
	"strings"
	"testing"
)

func TestDay01(t *testing.T) {
	sampleInput := `199
200
208
210
200
207
240
269
260
263`

	input := strings.Split(sampleInput, "\n")

	res := Day01(input)

	if expected := 7; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 5; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
