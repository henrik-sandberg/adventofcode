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

func day15_sovler(m map[complex64]rune, bot complex64, moves string) (res int) {
	directions := map[rune]complex64{'<': -1, '^': -1i, '>': 1, 'v': 1i}
moveloop:
	for _, move := range moves {
		dir := directions[move]
		pointer := bot
		tiles := []complex64{pointer}
		seen := []complex64{pointer}
		for len(tiles) > 0 {
			pointer, tiles = tiles[0], tiles[1:]
			if !slices.Contains(seen, pointer) {
				seen = append(seen, pointer)
			}
			pointer += dir
			switch m[pointer] {
			case '#':
				continue moveloop
			case '[':
				tiles = append(tiles, pointer)
				if imag(dir) != 0 {
					tiles = append(tiles, pointer+1)
				}
			case ']':
				tiles = append(tiles, pointer)
				if imag(dir) != 0 {
					tiles = append(tiles, pointer-1)
				}
			case 'O':
				tiles = append(tiles, pointer)
			}
		}
		slices.Reverse(seen)
		for _, tile := range seen {
			m[tile], m[tile+dir] = m[tile+dir], m[tile]
		}
		bot += dir
	}
	for k, v := range m {
		if v == 'O' || v == '[' {
			res += 100*int(imag(k)) + int(real(k))
		}
	}
	return
}
