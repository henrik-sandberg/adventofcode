package y2024

import (
	"adventofcode/shared"
	"strconv"
)

func Day21(input []string) shared.Solution {
	keypad := map[rune]map[rune]string{
		'A': {'<': "v<<A", '>': "vA", '^': "<A", 'v': "<vA", 'A': "A", '0': "<A", '1': "^<<A", '2': "<^A", '3': "^A", '4': "^^<<A", '5': "<^^A", '6': "^^A", '7': "^^^<<A", '8': "<^^^A", '9': "^^^A"},
		'^': {'<': "v<A", '>': "v>A", '^': "A", 'v': "vA", 'A': ">A"},
		'v': {'<': "<A", '>': ">A", '^': "^A", 'v': "A", 'A': "^>A"},
		'<': {'<': "A", '>': ">>A", '^': ">^A", 'v': ">A", 'A': ">>^A"},
		'>': {'<': "<<A", '>': "A", '^': "<^A", 'v': "<A", 'A': "^A"},
		'0': {'1': "^<A", '2': "^A", '3': "^>A", '4': "^<^A", '5': "^^A", '6': "^^>A", '7': "^^^<A", '8': "^^^A", '9': "^^^>A", 'A': ">A"},
		'1': {'0': ">vA", '2': ">A", '3': ">>A", '4': "^A", '5': "^>A", '6': "^>>A", '7': "^^A", '8': "^^>A", '9': "^^>>A", 'A': ">>vA"},
		'2': {'0': "vA", '1': "<A", '3': ">A", '4': "<^A", '5': "^A", '6': "^>A", '7': "<^^A", '8': "^^A", '9': "^^>A", 'A': "v>A"},
		'3': {'0': "<vA", '1': "<<A", '2': "<A", '4': "<<^A", '5': "<^A", '6': "^A", '7': "<<^^A", '8': "<^^A", '9': "^^A", 'A': "vA"},
		'4': {'0': ">vvA", '1': "vA", '2': "v>A", '3': "v>>A", '5': ">A", '6': ">>A", '7': "^A", '8': "^>A", '9': "^>>A", 'A': ">>vvA"},
		'5': {'0': "vvA", '1': "<vA", '2': "vA", '3': "v>A", '4': "<A", '6': ">A", '7': "<^A", '8': "^A", '9': "^>A", 'A': "vv>A"},
		'6': {'0': "<vvA", '1': "<<vA", '2': "<vA", '3': "vA", '4': "<<A", '5': "<A", '7': "<<^A", '8': "<^A", '9': "^A", 'A': "vvA"},
		'7': {'0': ">vvvA", '1': "vvA", '2': "vv>A", '3': "vv>>A", '4': "vA", '5': "v>A", '6': "v>>A", '8': ">A", '9': ">>A", 'A': ">>vvvA"},
		'8': {'0': "vvvA", '1': "<vvA", '2': "vvA", '3': "vv>A", '4': "<vA", '5': "vA", '6': "v>A", '7': "<A", '9': ">A", 'A': "vvv>A"},
		'9': {'0': "<vvvA", '1': "<<vvA", '2': "<vvA", '3': "vvA", '4': "<<vA", '5': "<vA", '6': "vA", '7': "<<A", '8': "<A", 'A': "vvvA"},
	}
	solver := func(maxdepth int) (res int) {
		type key struct {
			seq   string
			depth int
		}
		cache := map[key]int{}
		var expand func(string, int) int
		expand = func(seq string, depth int) (res int) {
			if depth == maxdepth {
				return len(seq)
			}
			k := key{seq, depth}
			if v, ok := cache[k]; ok {
				return v
			}
			current := 'A'
			for _, r := range seq {
				next := keypad[current][r]
				res += expand(next, depth+1)
				current = r
			}
			cache[k] = res
			return res
		}
		for _, code := range input {
			val, _ := strconv.Atoi(code[:len(code)-1])
			res += val * expand(code, 0)
		}
		return res
	}
	return shared.Solution{
		Part1: solver(3),
		Part2: solver(26),
	}
}
