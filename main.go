package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	problem := os.Args[1]
	filename := fmt.Sprintf("input/day%s.txt", problem)
	if len(os.Args) > 2 {
		filename = os.Args[2]
	}
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	input := strings.Split(strings.TrimSpace(string(content)), "\n")

	fmt.Printf("Running problem: %s with file: %s\n", problem, filename)
	switch problem {
	case "01":
		Day01(input)
	case "02":
		Day02(input)
	case "03":
		Day03(input)
	case "04":
		Day04(input)
	case "05":
		Day05(input)
	case "06":
		Day06(input)
	case "07":
		Day07(input)
	case "08":
		Day08(input)
	case "09":
		Day09(input)
	case "10":
		Day10(input)
	case "11":
		Day11(input)
	case "12":
		Day12(input)
	case "13":
		Day13(input)
	case "14":
		Day14(input)
	case "15":
		Day15(input)
	case "16":
		Day16(input)
	case "17":
		Day17(input)
	case "18":
		Day18(input)
	case "19":
		Day19(input)
	case "20":
		Day20(input)
	case "21":
		Day21(input)
	case "22":
		Day22(input)
	case "23":
		Day23(input)
	case "24":
		Day24(input)
	case "25":
		Day25(input)
	default:
		fmt.Printf("Problem %s not implemented\n", problem)
	}
}
