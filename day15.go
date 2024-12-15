package main

import (
	"fmt"
	"slices"
	"strings"
)

func Day15(input []string) {
	directions := map[rune]complex64{'<': -1, '^': -1i, '>': 1, 'v': 1i}
	splitindex := slices.Index(input, "")
	flatmap := Map{strings.Join(input[:splitindex], ""), len(input[0])}
	m := flatmap.ToComplexGrid()
	bot := flatmap.GetPointComplex(strings.Index(flatmap.cells, "@"))
	m[bot] = '.'
	for _, move := range strings.Join(input[splitindex:], "") {
		dir := directions[move]
		// Find next empty space
		steps := 1
		pointer := bot + dir
		for m[pointer] != '#' && m[pointer] != '.' {
			pointer += dir
			steps++
		}
		fmt.Println("Found", string(m[pointer]), "at distance", steps)
		if m[pointer] == '#' {
			fmt.Println("No free space found in direction", dir)
			continue
		}
		for i := 0; i < steps; i++ {
			fmt.Println("Moving", string(m[pointer-dir]), "from", pointer-dir, "to", pointer)
			m[pointer] = m[pointer-dir]
			pointer -= dir
		}
		fmt.Println("Moving bot to", bot+dir)
		bot += dir
	}
	part1 := 0
	for k, v := range m {
		if v == 'O' {
			part1 += 100*int(imag(k)) + int(real(k))
		}
	}
	fmt.Println(part1)
	/*
		for r := 0; r < 8; r++ {
			sb := strings.Builder{}
			for c := 0; c < 8; c++ {
				cn := complex(float32(c), float32(r))
				if cn == bot {
					sb.WriteRune('@')
				} else {
					sb.WriteRune(m[cn])
				}
			}
			fmt.Println(sb.String())
		}
	*/
}
