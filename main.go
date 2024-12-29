package main

import (
	"adventofcode/y2024"
	"os"
)

func main() {
	year := os.Args[1]
	day := os.Args[2]
	switch year {
	case "2024":
		y2024.Run(day)
	}
}
