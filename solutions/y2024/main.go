package y2024

import (
	"adventofcode/shared"
)

// Solutions maps day numbers to their respective solution functions.
// Each function takes []string input and returns a shared.Solution.
var Solutions = map[string]func([]string) shared.Solution[any, any]{
	"01": func(input []string) shared.Solution[any, any] {
		sol := Day01(input)
		return shared.Solution[any, any]{Part1: sol.Part1, Part2: sol.Part2}
	},
	"02": func(input []string) shared.Solution[any, any] {
		sol := Day02(input)
		return shared.Solution[any, any]{Part1: sol.Part1, Part2: sol.Part2}
	},
	"03": func(input []string) shared.Solution[any, any] {
		sol := Day03(input)
		return shared.Solution[any, any]{Part1: sol.Part1, Part2: sol.Part2}
	},
	"04": func(input []string) shared.Solution[any, any] {
		sol := Day04(input)
		return shared.Solution[any, any]{Part1: sol.Part1, Part2: sol.Part2}
	},
	"05": func(input []string) shared.Solution[any, any] {
		sol := Day05(input)
		return shared.Solution[any, any]{Part1: sol.Part1, Part2: sol.Part2}
	},
	"06": func(input []string) shared.Solution[any, any] {
		sol := Day06(input)
		return shared.Solution[any, any]{Part1: sol.Part1, Part2: sol.Part2}
	},
	"07": func(input []string) shared.Solution[any, any] {
		sol := Day07(input)
		return shared.Solution[any, any]{Part1: sol.Part1, Part2: sol.Part2}
	},
	"08": func(input []string) shared.Solution[any, any] {
		sol := Day08(input)
		return shared.Solution[any, any]{Part1: sol.Part1, Part2: sol.Part2}
	},
	"09": func(input []string) shared.Solution[any, any] {
		sol := Day09(input)
		return shared.Solution[any, any]{Part1: sol.Part1, Part2: sol.Part2}
	},
	"10": func(input []string) shared.Solution[any, any] {
		sol := Day10(input)
		return shared.Solution[any, any]{Part1: sol.Part1, Part2: sol.Part2}
	},
	"11": func(input []string) shared.Solution[any, any] {
		sol := Day11(input)
		return shared.Solution[any, any]{Part1: sol.Part1, Part2: sol.Part2}
	},
	"12": func(input []string) shared.Solution[any, any] {
		sol := Day12(input)
		return shared.Solution[any, any]{Part1: sol.Part1, Part2: sol.Part2}
	},
	"13": func(input []string) shared.Solution[any, any] {
		sol := Day13(input)
		return shared.Solution[any, any]{Part1: sol.Part1, Part2: sol.Part2}
	},
	"14": func(input []string) shared.Solution[any, any] {
		sol := Day14(input)
		return shared.Solution[any, any]{Part1: sol.Part1, Part2: sol.Part2}
	},
	"15": func(input []string) shared.Solution[any, any] {
		sol := Day15(input)
		return shared.Solution[any, any]{Part1: sol.Part1, Part2: sol.Part2}
	},
	"16": func(input []string) shared.Solution[any, any] {
		sol := Day16(input)
		return shared.Solution[any, any]{Part1: sol.Part1, Part2: sol.Part2}
	},
	"17": func(input []string) shared.Solution[any, any] {
		sol := Day17(input)
		return shared.Solution[any, any]{Part1: sol.Part1, Part2: sol.Part2}
	},
	"18": func(input []string) shared.Solution[any, any] {
		sol := Day18(input)
		return shared.Solution[any, any]{Part1: sol.Part1, Part2: sol.Part2}
	},
	"19": func(input []string) shared.Solution[any, any] {
		sol := Day19(input)
		return shared.Solution[any, any]{Part1: sol.Part1, Part2: sol.Part2}
	},
	"20": func(input []string) shared.Solution[any, any] {
		sol := Day20(input)
		return shared.Solution[any, any]{Part1: sol.Part1, Part2: sol.Part2}
	},
	"21": func(input []string) shared.Solution[any, any] {
		sol := Day21(input)
		return shared.Solution[any, any]{Part1: sol.Part1, Part2: sol.Part2}
	},
	"22": func(input []string) shared.Solution[any, any] {
		sol := Day22(input)
		return shared.Solution[any, any]{Part1: sol.Part1, Part2: sol.Part2}
	},
	"23": func(input []string) shared.Solution[any, any] {
		sol := Day23(input)
		return shared.Solution[any, any]{Part1: sol.Part1, Part2: sol.Part2}
	},
	"24": func(input []string) shared.Solution[any, any] {
		sol := Day24(input)
		return shared.Solution[any, any]{Part1: sol.Part1, Part2: sol.Part2}
	},
	"25": func(input []string) shared.Solution[any, any] {
		sol := Day25(input)
		return shared.Solution[any, any]{Part1: sol.Part1, Part2: sol.Part2}
	},
}

