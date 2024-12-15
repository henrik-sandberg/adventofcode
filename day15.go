package main

import (
	"fmt"
	"slices"
	"strings"
)

func Day15(input []string) {
	directions := map[rune]complex64{'<': -1, '^': -1i, '>': 1, 'v': 1i}
	splitindex := slices.Index(input, "")

	/*
		for i := 0; i < splitindex; i++ {
			input[i] = strings.ReplaceAll(input[i], "#", "##")
			input[i] = strings.ReplaceAll(input[i], "O", "[]")
			input[i] = strings.ReplaceAll(input[i], ".", "..")
			input[i] = strings.ReplaceAll(input[i], "@", "@.")
		}
	*/
	flatmap := Map{strings.Join(input[:splitindex], ""), len(input[0])}
	m := flatmap.ToComplexGrid()
	bot := flatmap.GetPointComplex(strings.Index(flatmap.cells, "@"))
	m[bot] = '.'
	printmap := func() {
		for r := 0; r < flatmap.Height(); r++ {
			sb := strings.Builder{}
			for c := 0; c < flatmap.Width(); c++ {
				cn := complex(float32(c), float32(r))
				if cn == bot {
					sb.WriteRune('@')
				} else {
					sb.WriteRune(m[cn])
				}
			}
			// fmt.Println(sb.String())
		}
	}
	printmap()
moveloop:
	for _, move := range strings.Join(input[splitindex:], "")[:] {
		dir := directions[move]
		steps := 1
		pointer := bot + dir
		if m[pointer] == '#' {
			continue
		}
		if m[pointer] == '.' {
			bot += dir
			continue
		}

		if imag(dir) != 0 {
			box := pointer
			boxes := []complex64{box}

			switch m[pointer] {
			case '[':
				boxes = append(boxes, box+1)
			case ']':
				boxes = append(boxes, box-1)
			}
			traversed := []complex64{}
			for len(boxes) > 0 {
				box, boxes = boxes[0], boxes[1:]
				next := box + dir
				if !slices.Contains(traversed, box) {
					traversed = append(traversed, box)
				}
				switch m[next] {
				case '#':
					continue moveloop
				case '[':
					boxes = append(boxes, next, next+1)
				case ']':
					boxes = append(boxes, next, next-1)
				case 'O':
					boxes = append(boxes, next)
				}
				if v := abs(int(real(box)) - int(real(bot))); v > steps {
					steps = v
				}
			}
			slices.Reverse(traversed)
			for _, box := range traversed {
				m[box], m[box+dir] = m[box+dir], m[box]
			}
			bot += dir
			continue
		}

		for m[pointer] != '#' && m[pointer] != '.' {
			pointer += dir
			steps++
		}
		if m[pointer] == '#' {
			continue
		}
		for i := 0; i < steps; i++ {
			m[pointer] = m[pointer-dir]
			pointer -= dir
		}
		bot += dir
	}
	part2 := 0
	for k, v := range m {
		if v == 'O' {
			part2 += 100*int(imag(k)) + int(real(k))
		}
	}
	fmt.Println(part2)

	// PART 1

	for _, move := range strings.Join(input[splitindex:], "")[:0] {
		dir := directions[move]
		// Find next empty space
		steps := 1
		pointer := bot + dir
		for m[pointer] != '#' && m[pointer] != '.' {
			pointer += dir
			steps++
		}
		if m[pointer] == '#' {
			continue
		}
		for i := 0; i < steps; i++ {
			m[pointer] = m[pointer-dir]
			pointer -= dir
		}
		bot += dir
	}
	part1 := 0
	for k, v := range m {
		if v == 'O' {
			part1 += 100*int(imag(k)) + int(real(k))
		}
	}
	fmt.Println(part1)
}
