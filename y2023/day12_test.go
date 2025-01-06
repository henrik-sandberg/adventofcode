package y2023

import (
	"strings"
	"testing"
)

func TestDay12(t *testing.T) {
	sampleInput := strings.Split(strings.TrimSpace(`
???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1
	`), "\n")

	res := Day12(sampleInput)

	if expected := 21; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 525152; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
