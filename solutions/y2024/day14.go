package y2024

import (
	"adventofcode/shared"
	"math"
	"regexp"
)

func Day14(input []string) (solution shared.Solution[int, int]) {
	re := regexp.MustCompile(`-?\d+`)
	bots := make([][]int, len(input))
	for i, line := range input {
		bots[i] = shared.IntSlice(re.FindAllString(line, -1))
	}
	width := 101
	height := 103
	move := func(bot []int, rounds int) (int, int) {
		x := shared.PositiveMod(bot[0]+bot[2]*rounds, width)
		y := shared.PositiveMod(bot[1]+bot[3]*rounds, height)
		return x, y
	}
	quadrants := []int{0, 0, 0, 0}
	for _, bot := range bots {
		if x, y := move(bot, 100); x != width/2 && y != height/2 {
			quadrants[2*x/width+2*(2*y/height)]++
		}
	}
	solution.Part1 = shared.Product(quadrants...)
	mindev := math.MaxFloat32
	for n := 1; n < 10_000; n++ {
		X, Y := make([]int, len(bots)), make([]int, len(bots))
		for i, bot := range bots {
			X[i], Y[i] = move(bot, n)
		}
		if dev := shared.Stddev(X) + shared.Stddev(Y); dev < mindev {
			mindev = dev
			solution.Part2 = n
		}
	}
	return
}
