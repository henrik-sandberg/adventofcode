package shared

import (
	"strconv"
	"strings"
)

func Reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func IntSlice(input []string) []int {
	ints := make([]int, len(input))
	for i, s := range input {
		ints[i], _ = strconv.Atoi(strings.TrimSpace(s))
	}
	return ints
}
