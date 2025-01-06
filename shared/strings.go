package shared

import "strconv"

func Reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func IntSlice(input []string) []int {
	ints := make([]int, len(input), len(input))
	for i, s := range input {
		value, _ := strconv.Atoi(s)
		ints[i] = value
	}
	return ints
}
