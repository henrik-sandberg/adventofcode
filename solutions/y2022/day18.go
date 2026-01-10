package y2022

import (
	"adventofcode/solutions/shared"
	"iter"
	"strconv"
	"strings"
)

func Day18(input []string) (solution shared.Solution[int, int]) {
	type cube struct {
		x, y, z int
	}
	sides := func(c cube) iter.Seq[cube] {
		cubeDirs := []cube{
			{1, 0, 0}, {-1, 0, 0},
			{0, 1, 0}, {0, -1, 0},
			{0, 0, 1}, {0, 0, -1},
		}
		return func(yield func(cube) bool) {
			for _, d := range cubeDirs {
				if !yield(cube{
					x: c.x + d.x,
					y: c.y + d.y,
					z: c.z + d.z,
				}) {
					return
				}
			}
		}
	}
	cubes := make(map[cube]bool)
	for _, line := range input {
		c := strings.Split(line, ",")
		x, _ := strconv.Atoi(c[0])
		y, _ := strconv.Atoi(c[1])
		z, _ := strconv.Atoi(c[2])
		cubes[cube{x, y, z}] = true
	}
	for c := range cubes {
		for side := range sides(c) {
			if !cubes[side] {
				solution.Part1++
			}
		}
	}
	var c cube
	queue := []cube{{-1, -1, -1}}
	seen := make(map[cube]bool)
	for len(queue) > 0 {
		c, queue = queue[0], queue[1:]
		for side := range sides(c) {
			if cubes[side] {
				solution.Part2++
			} else if !seen[side] &&
				c.x >= -1 && c.x < 21 &&
				c.y >= -1 && c.y < 21 &&
				c.z >= -1 && c.z < 21 {
				seen[side] = true
				queue = append(queue, side)
			}
		}
	}
	return
}
