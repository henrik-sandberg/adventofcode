package y2024

import (
	"slices"
	"strings"

	"adventofcode/solutions/shared"
)

func Day15(input []string) (solution shared.Solution[int, int]) {
	splitindex := slices.Index(input, "")
	moves := strings.Join(input[splitindex:], "")
	setup := func() (shared.Grid, complex128) {
		grid := shared.NewGrid(input)
		bot := grid.FindAny('@')
		grid[bot] = '.'
		return grid, bot
	}
	m, bot := setup()
	solution.Part1 = day15_sovler(m, bot, moves)
	for i := range splitindex {
		input[i] = strings.ReplaceAll(input[i], "#", "##")
		input[i] = strings.ReplaceAll(input[i], "O", "[]")
		input[i] = strings.ReplaceAll(input[i], ".", "..")
		input[i] = strings.ReplaceAll(input[i], "@", "@.")
	}
	m, bot = setup()
	solution.Part2 = day15_sovler(m, bot, moves)
	return
}

func day15_sovler(m shared.Grid, bot complex128, moves string) (res int) {
	directions := map[rune]complex128{'<': -1, '^': -1i, '>': 1, 'v': 1i}
moveloop:
	for _, move := range moves {
		dir := directions[move]
		pointer := bot
		tiles := []complex128{pointer}
		seen := []complex128{pointer}
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
