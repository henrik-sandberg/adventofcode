package y2025

import (
	"adventofcode/solutions/shared"
	"slices"
	"strings"
)

func Day09(input []string) (solution shared.Solution[int, int]) {
	type point struct {
		x, y int
	}
	type bounds struct {
		left, right, top, bottom int
	}
	makeBounds := func(a, b point) bounds {
		return bounds{
			left:   min(a.x, b.x),
			right:  max(a.x, b.x),
			top:    min(a.y, b.y),
			bottom: max(a.y, b.y),
		}
	}
	area := func(b bounds) int {
		return (b.right - b.left + 1) * (b.bottom - b.top + 1)
	}
	points := make([]point, len(input))
	for i, s := range input {
		ints := shared.IntSlice(strings.Split(s, ","))
		points[i] = point{x: ints[0], y: ints[1]}
	}
	var rectangles []bounds
	for ps := range shared.Combinations(points, 2) {
		rectangles = append(rectangles, makeBounds(ps[0], ps[1]))
	}
	slices.SortFunc(rectangles, func(a, b bounds) int {
		return area(b) - area(a)
	})
	solution.Part1 = area(rectangles[0])
	intersectsPolyline := func(rb bounds) bool {
		for i := range points {
			pb := makeBounds(points[i], points[(i+1)%len(points)])
			if pb.left < rb.right &&
				pb.right > rb.left &&
				pb.bottom > rb.top &&
				pb.top < rb.bottom {
				return true
			}
		}
		return false
	}
	for _, r := range rectangles {
		if !intersectsPolyline(r) {
			solution.Part2 = area(r)
			break
		}
	}
	return
}
