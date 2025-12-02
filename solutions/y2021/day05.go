package y2021

import (
	"adventofcode/solutions/shared"
	"maps"
	"slices"
	"strings"
)

func Day05(input []string) (solution shared.Solution[int, int]) {
	type point struct {
		x, y int
	}
	type segment struct {
		from, to point
	}
	var segments []segment
	for _, s := range input {
		cords := strings.Split(s, " -> ")
		from := shared.IntSlice(strings.Split(cords[0], ","))
		to := shared.IntSlice(strings.Split(cords[1], ","))
		segments = append(segments, segment{
			from: point{x: from[0], y: from[1]},
			to:   point{x: to[0], y: to[1]},
		})
	}

	sign := func(i int) int {
		if i > 0 {
			return 1
		}
		if i < 0 {
			return -1
		}
		return 0
	}
	overlaps := func(segments []segment) int {
		grid := map[point]int{}
		for _, seg := range segments {
			dx := sign(seg.to.x - seg.from.x)
			dy := sign(seg.to.y - seg.from.y)
			x, y := seg.from.x, seg.from.y
			for {
				p := point{x: x, y: y}
				grid[p]++
				if p == seg.to {
					break
				}
				x += dx
				y += dy
			}
		}
		ret := 0
		for val := range maps.Values(grid) {
			if val > 1 {
				ret++
			}
		}
		return ret
	}
	solution.Part2 = overlaps(segments)
	segments = slices.DeleteFunc(segments, func(s segment) bool {
		return s.from.x != s.to.x && s.from.y != s.to.y
	})
	solution.Part1 = overlaps(segments)
	return
}
