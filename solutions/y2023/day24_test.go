package y2023

import (
	"strings"
	"testing"
)

func TestDay24(t *testing.T) {
	sampleInput := strings.Split(strings.TrimSpace(`
19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3
	`), "\n")

	day24 := day24{sampleInput}
	res := day24.part1(7, 27)
	if expected := 2; expected != res {
		t.Errorf("part1: expected %d, got %d", expected, res)
	}
}
