package y2022

import (
	"adventofcode/solutions/shared"
	"maps"
	"math"
	"regexp"
	"strconv"
)

func Day15(input []string) shared.Solution[int, int] {
	return day15solve(input, 2000000)
}

func day15solve(input []string, row int) (solution shared.Solution[int, int]) {
	type Point struct {
		x, y int
	}
	manhattan := func(a, b Point) int {
		return shared.Abs(a.x-b.x) + shared.Abs(a.y-b.y)
	}
	re := regexp.MustCompile(`[xy]=(-?\d+)`)
	parse := func(match [][]string, i int) Point {
		x, _ := strconv.Atoi(match[i][1])
		y, _ := strconv.Atoi(match[i+1][1])
		return Point{x: x, y: y}
	}
	radius := make(map[Point]int)
	beacons := make(map[Point]struct{})
	for _, line := range input {
		match := re.FindAllStringSubmatch(line, 4)
		sensor := parse(match, 0)
		beacon := parse(match, 2)
		beacons[beacon] = struct{}{}
		radius[sensor] = manhattan(sensor, beacon)
	}
	xmin, xmax := math.MaxInt, math.MinInt
	for p, dist := range radius {
		remaining := dist - shared.Abs(p.y-row)
		xmin = min(xmin, p.x-remaining)
		xmax = max(xmax, p.x+remaining)
	}
	maps.DeleteFunc(beacons, func(p Point, _ struct{}) bool {
		return p.y != row
	})
	solution.Part1 = xmax - xmin + 1 - len(beacons)

	findNonCoveredPoint := func() Point {
		acoeffs := make(map[int]struct{})
		bcoeffs := make(map[int]struct{})
		for p, r := range radius {
			acoeffs[p.y-p.x+r+1] = struct{}{}
			acoeffs[p.y-p.x-r-1] = struct{}{}
			bcoeffs[p.x+p.y+r+1] = struct{}{}
			bcoeffs[p.x+p.y-r-1] = struct{}{}
		}
		bound := 2 * row
		for a := range acoeffs {
			for b := range bcoeffs {
				p := Point{x: (b - a) >> 1, y: (a + b) >> 1}
				if p.x < 0 || p.x > bound ||
					p.y < 0 || p.y > bound {
					continue
				}
				valid := true
				for s, dist := range radius {
					if manhattan(s, p) <= dist {
						valid = false
						break
					}
				}
				if valid {
					return p
				}
			}
		}
		panic("could not find non/covered point")
	}
	p := findNonCoveredPoint()
	solution.Part2 = p.x*4000000 + p.y
	return
}
