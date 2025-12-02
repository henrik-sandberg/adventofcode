package y2023

import (
	"slices"
	"strings"

	"adventofcode/solutions/shared"
)

type lens struct {
	label string
	focal int
}

func Day15(input []string) (solution shared.Solution[int, int]) {
	for s := range strings.SplitSeq(input[0], ",") {
		solution.Part1 += hash(s)
	}
	boxes := make([][]lens, 256)
	for s := range strings.SplitSeq(input[0], ",") {
		if s[len(s)-1] == '-' {
			label := s[:len(s)-1]
			h := hash(label)
			if ind := slices.IndexFunc(boxes[h], func(l lens) bool {
				return l.label == label
			}); ind != -1 {
				boxes[h] = Remove(boxes[h], ind)
			}
		} else {
			label := s[:len(s)-2]
			h := hash(label)
			lab := lens{label, int(s[len(s)-1] - '0')}
			if ind := slices.IndexFunc(boxes[h], func(l lens) bool {
				return l.label == label
			}); ind != -1 {
				boxes[h][ind] = lab
			} else {
				boxes[h] = append(boxes[h], lab)
			}
		}
	}
	for bi, box := range boxes {
		for slot, l := range box {
			solution.Part2 += (bi + 1) * (slot + 1) * l.focal
		}
	}
	return
}

func hash(s string) int {
	h := 0
	for _, c := range s {
		h = (h + int(c)) * 17 & 0xFF
	}
	return h
}

func Remove[T any](slice []T, index int) []T {
	ret := make([]T, 0, len(slice)-1)
	ret = append(ret, slice[:index]...)
	return append(ret, slice[index+1:]...)
}
