package y2023

import (
	"strings"
	"testing"
)

func TestDay15(t *testing.T) {
	sampleInput := "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"

	input := strings.Split(sampleInput, "\n")

	res := Day15(input)

	if expected := 1320; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 145; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
