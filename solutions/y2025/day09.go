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
	type rectangle struct {
		bounds
		area int
	}
	points := make([]point, len(input))
	for i, s := range input {
		ints := shared.IntSlice(strings.Split(s, ","))
		points[i] = point{x: ints[0], y: ints[1]}
	}
	makeBounds := func(a, b point) bounds {
		return bounds{
			left:   min(a.x, b.x),
			right:  max(a.x, b.x),
			top:    min(a.y, b.y),
			bottom: max(a.y, b.y),
		}
	}
	var rectangles []rectangle
	for ps := range shared.Combinations(points, 2) {
		bnds := makeBounds(ps[0], ps[1])
		area := (bnds.right - bnds.left + 1) * (bnds.bottom - bnds.top + 1)
		rectangles = append(rectangles, rectangle{bounds: bnds, area: area})
	}
	slices.SortFunc(rectangles, func(a, b rectangle) int {
		return b.area - a.area
	})
	solution.Part1 = rectangles[0].area
	intersectsPolyline := func(r rectangle) bool {
		for i := range points {
			pb := makeBounds(points[i], points[(i+1)%len(points)])
			if pb.left < r.right &&
				pb.right > r.left &&
				pb.bottom > r.top &&
				pb.top < r.bottom {
				return true
			}
		}
		return false
	}
	for _, r := range rectangles {
		if !intersectsPolyline(r) {
			solution.Part2 = r.area
			break
		}
	}
	return
}
