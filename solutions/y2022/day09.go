package y2022

import (
	"adventofcode/solutions/shared"
	"strconv"
	"strings"
)

func Day09(input []string) (solution shared.Solution[int, int]) {
	solution.Part1 = solveDay09(input, 2)
	solution.Part1 = solveDay09(input, 10)
	return
}

func solveDay09(input []string, length int) int {
	seen := []point{}
	rope := make([]point, length)
	for _, cmd := range input {
		cmdArr := strings.Split(cmd, " ")
		direction := cmdArr[0]
		distance, _ := strconv.Atoi(cmdArr[1])
		for n := 0; n < distance; n++ {
			rope[0].move(direction)
			for i := 1; i < len(rope); i++ {
				rope[i].follow(rope[i-1])
			}
			seen = append(seen, rope[len(rope)-1])
		}
	}
	unique := map[point]bool{}
	for _, p := range seen {
		unique[p] = true
	}
	return len(unique)
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
