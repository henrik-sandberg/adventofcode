package y2022

import (
	"strconv"
	"strings"

	"adventofcode/solutions/shared"
)

func Day09(input []string) (solution shared.Solution[int, int]) {
	solution.Part1 = solveDay09(input, 2)
	solution.Part2 = solveDay09(input, 10)
	return
}

func solveDay09(input []string, length int) int {
	var path []point
	rope := make([]point, length)
	for _, line := range input {
		direction, distStr, _ := strings.Cut(line, " ")
		dist, _ := strconv.Atoi(distStr)
		for range dist {
			rope[0].move(direction)
			for i := 1; i < len(rope); i++ {
				rope[i].follow(rope[i-1])
			}
			path = append(path, rope[len(rope)-1])
		}
	}
	seen := make(map[point]struct{})
	for _, p := range path {
		seen[p] = struct{}{}
	}
	return len(seen)
}

func (p *point) move(direction string) {
	switch direction {
	case "U":
		p.y += 1
	case "D":
		p.y -= 1
	case "R":
		p.x += 1
	case "L":
		p.x -= 1
	}
}

func (p *point) follow(head point) {
	if abs(head.x-p.x) > 1 || abs(head.y-p.y) > 1 {
		if head.x > p.x {
			p.x += 1
		} else if head.x < p.x {
			p.x -= 1
		}
		if head.y > p.y {
			p.y += 1
		} else if head.y < p.y {
			p.y -= 1
		}
	}
}
