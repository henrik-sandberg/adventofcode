package shared

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type Solution struct {
	Part1 any
	Part2 any
}

func Run(f func([]string) Solution) {
	input := readLines()
	start := time.Now()
	solution := f(input)
	elapsed := time.Since(start)

	fmt.Printf("Part1: %v\n", solution.Part1)
	fmt.Printf("Part2: %v\n", solution.Part2)
	fmt.Printf("Found solution in %d ms\n", elapsed.Milliseconds())
}

func readLines() []string {
	lines := []string{}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading from stdin:", err)
	}
	return lines
}