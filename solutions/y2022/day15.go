package y2022

import (
	"adventofcode/solutions/shared"
	"maps"
	"regexp"
	"slices"
	"strconv"
)

func Day15(input []string) shared.Solution[int, int] {
	return day15solve(input, 2000000)
}

func day15solve(input []string, row int) (solution shared.Solution[int, int]) {
	type Range struct {
		from, to int
	}
	type Point struct {
		x, y int
	}
	parse := func(match [][]string, i int) Point {
		x, _ := strconv.Atoi(match[i][1])
		y, _ := strconv.Atoi(match[i+1][1])
		return Point{x: x, y: y}
	}
	re := regexp.MustCompile(`[xy]=(-?\d+)`)
	sensors := make(map[Point]Point, len(input))
	beaconsOnTargetRow := make(map[Point]struct{})
	for _, line := range input {
		match := re.FindAllStringSubmatch(line, 4)
		sensor := parse(match, 0)
		beacon := parse(match, 2)
		sensors[sensor] = beacon
		beaconsOnTargetRow[beacon] = struct{}{}
	}
	maps.DeleteFunc(beaconsOnTargetRow, func(p Point, _ struct{}) bool {
		return p.y != row
	})
	coveredPoints := func(row int) []Range {
		var ranges []Range
		for sensor, beacon := range sensors {
			dist := shared.Abs(sensor.x-beacon.x) + shared.Abs(sensor.y-beacon.y)
			perpendicularToY := shared.Abs(sensor.y - row)
			if dist > perpendicularToY {
				remaining := dist - perpendicularToY
				ranges = append(ranges, Range{sensor.x - remaining, sensor.x + remaining})
			}
		}
		slices.SortFunc(ranges, func(a, b Range) int {
			return a.from - b.from
		})
		merged := []Range{ranges[0]}
		for _, r := range ranges[1:] {
			last := &merged[len(merged)-1]
			if r.from > last.to+1 {
				merged = append(merged, r)
			} else if r.to > last.to {
				last.to = r.to
			}
		}
		return merged
	}

	covered := coveredPoints(row)[0]
	solution.Part1 = covered.to - covered.from + 1 - len(beaconsOnTargetRow)

	for i := range 2*row + 1 {
		covered := coveredPoints(i)
		if len(covered) > 1 {
			solution.Part2 = (covered[0].to+1)*4000000 + i
			break
		}
	}
	return
}
