package main

import (
	"fmt"
	"slices"
	"strings"
)

func Day15(input []string) {
	splitindex := slices.Index(input, "")
	moves := strings.Join(input[splitindex:], "")
	setup := func() (map[complex64]rune, complex64) {
		inpmap := Map{strings.Join(input[:splitindex], ""), len(input[0])}
		bot := inpmap.GetPointComplex(strings.Index(inpmap.cells, "@"))
		m := inpmap.ToComplexGrid()
		m[bot] = '.'
		return m, bot
	}

	m, bot := setup()
	fmt.Println(day15_sovler(m, bot, moves))
	for i := 0; i < splitindex; i++ {
		input[i] = strings.ReplaceAll(input[i], "#", "##")
		input[i] = strings.ReplaceAll(input[i], "O", "[]")
		input[i] = strings.ReplaceAll(input[i], ".", "..")
		input[i] = strings.ReplaceAll(input[i], "@", "@.")
	}
	m, bot = setup()
	fmt.Println(day15_sovler(m, bot, moves))
}

func day15_sovler(m map[complex64]rune, bot complex64, moves string) int {
	directions := map[rune]complex64{'<': -1, '^': -1i, '>': 1, 'v': 1i}
moveloop:
	for _, move := range moves {
		dir := directions[move]
		pointer := bot + dir
		if m[pointer] == '#' {
			continue
		}
		if m[pointer] == '.' {
			bot += dir
			continue
		}
		box := pointer
		boxes := []complex64{box}
		if imag(dir) != 0 {
			switch m[pointer] {
			case '[':
				boxes = append(boxes, box+1)
			case ']':
				boxes = append(boxes, box-1)
			}
		}
		seen := []complex64{}
		for len(boxes) > 0 {
			box, boxes = boxes[0], boxes[1:]
			if !slices.Contains(seen, box) {
				seen = append(seen, box)
			}
			next := box + dir
			switch m[next] {
			case '#':
				continue moveloop
			case '[':
				boxes = append(boxes, next)
				if imag(dir) != 0 {
					boxes = append(boxes, next+1)
				}
			case ']':
				boxes = append(boxes, next)
				if imag(dir) != 0 {
					boxes = append(boxes, next-1)
				}
			case 'O':
				boxes = append(boxes, next)
			}
		}
		slices.Reverse(seen)
		for _, box := range seen {
			m[box], m[box+dir] = m[box+dir], m[box]
		}
		bot += dir
	}
	res := 0
	for k, v := range m {
		if v == 'O' || v == '[' {
			res += 100*int(imag(k)) + int(real(k))
		}
	}
	return res
}
