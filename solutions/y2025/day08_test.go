package y2025

import (
	"strings"
	"testing"
)

func TestDay08(t *testing.T) {
	sampleInput := `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`

	input := strings.Split(sampleInput, "\n")

	res := (&day08{}).solve(input, 10)

	if expected := 40; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 25272; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
