package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day03(input []string) {
	inp := strings.Join(input, "")
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|(do|don't)\(\)`)
	part1 := 0
	part2 := 0
	enabled := true
	for _, match := range re.FindAllStringSubmatch(inp, -1) {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		part1 += a * b
		if match[3] == "do" {
			enabled = true
		} else if match[3] == "don't" {
			enabled = false
		}
		if enabled {
			part2 += a * b
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
}
