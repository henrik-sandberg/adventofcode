package main

import (
	"fmt"
	"regexp"
	"strings"
)

func Day14(input []string) {
	re := regexp.MustCompile(`-?\d+`)
	bots := make([][]int, len(input))
	for i, line := range input {
		bots[i] = ToInts(re.FindAllString(line, -1))
	}
	width := 101
	height := 103
	move := func(bot []int, rounds int) (int, int) {
		x := positiveMod(bot[0]+bot[2]*rounds, width)
		y := positiveMod(bot[1]+bot[3]*rounds, height)
		return x, y
	}
	quadrants := []int{0, 0, 0, 0}
	for _, bot := range bots {
		if x, y := move(bot, 100); x != width/2 && y != height/2 {
			quadrants[2*x/width+2*(2*y/height)]++
		}
	}
	fmt.Println(multiply(quadrants...))
	for n := 1; ; n++ {
		X, Y := make([]int, len(bots)), make([]int, len(bots))
		for i, bot := range bots {
			X[i], Y[i] = move(bot, n)
		}
		if sigx, sigy := stddev(X), stddev(Y); sigx+sigy < 40 {
			fmt.Println(n)
			break
		}
	}
}

func printChristmasTree(botX, botY []int, width, height int) {
	output := make([][]string, height)
	for h := 0; h < height; h++ {
		row := make([]string, width)
		for w := 0; w < width; w++ {
			row[w] = "."
		}
		output[h] = row
	}
	for i := range botX {
		output[botY[i]][botX[i]] = "x"
	}
	for _, line := range output {
		fmt.Println(strings.Join(line, ""))
	}
}
