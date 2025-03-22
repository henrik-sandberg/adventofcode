package y2023

import (
	"strings"
	"testing"
)

func TestDay22(t *testing.T) {
	sampleInput := strings.Split(strings.TrimSpace(`
1,0,1~1,2,1
0,0,2~2,0,2
0,2,3~2,2,3
0,0,4~0,2,4
2,0,5~2,2,5
0,1,6~2,1,6
1,1,8~1,1,9
`), "\n")

	res := Day22(sampleInput)

	if expected := 5; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 7; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
