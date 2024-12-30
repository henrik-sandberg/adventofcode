package main

import (
	"adventofcode/y2024"
	"adventofcode/y2025"
	"os"
)

func main() {
	year := os.Args[1]
	day := os.Args[2]
	switch year {
	case "2024":
		y2024.Run(day)
	case "2025":
		y2025.Run(day)
	}
}
