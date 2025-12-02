package y2023

import (
	"math"
	"strconv"
	"strings"

	"adventofcode/solutions/shared"
)

func Day18(input []string) (solution shared.Solution[int, int]) {
	move := func(vs []complex128, direction string, steps int) complex128 {
		var dir complex128
		switch direction {
		case "R", "0":
			dir = 1
		case "D", "1":
			dir = 1i
		case "L", "2":
			dir = -1
		case "U", "3":
			dir = -1i
		}
		return vs[len(vs)-1] + dir*complex(float64(steps), 0)
	}
	solver := func(vs []complex128) int {
		innerArea := float64(0)
		perimeter := 0
		for i := range len(vs) - 1 {
			a, b := vs[i], vs[i+1]
			innerArea += real(a)*imag(b) - real(b)*imag(a)
			perimeter += int(math.Abs(real(a)-real(b)) + math.Abs(imag(a)-imag(b)))
		}
		return int(innerArea/2) + perimeter/2 + 1
	}
	vertices := []complex128{0}
	for _, line := range input {
		fields := strings.Fields(line)
		steps, _ := strconv.Atoi(fields[1])
		vertices = append(vertices, move(vertices, fields[0], steps))
	}
	solution.Part1 = solver(vertices)
	vertices = []complex128{0}
	for _, line := range input {
		hex := strings.Fields(line)[2][2:8]
		steps, _ := strconv.ParseInt(hex[:5], 16, 0)
		vertices = append(vertices, move(vertices, hex[5:], int(steps)))
	}
	solution.Part2 = solver(vertices)
	return
}
