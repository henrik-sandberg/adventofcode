package y2023

import (
	"adventofcode/shared"
	"log"
)

func Run(day string) {
	switch day {
	case "01":
		shared.Run(Day01)
	case "02":
		shared.Run(Day02)
	case "03":
		shared.Run(Day03)
	case "04":
		shared.Run(Day04)
	case "05":
		shared.Run(Day05)
	case "06":
		shared.Run(Day06)
	case "07":
		shared.Run(Day07)
	case "08":
		shared.Run(Day08)
	case "09":
		shared.Run(Day09)
	case "10":
		shared.Run(Day10)
	case "11":
		shared.Run(Day11)
	case "12":
		shared.Run(Day12)
	case "13":
		shared.Run(Day13)
	case "14":
		shared.Run(Day14)
	case "15":
		shared.Run(Day15)
	default:
		log.Fatalln("Solution for day", day, "not implemented.")
	}
}
