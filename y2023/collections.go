package y2023

func bag(s string) map[rune]int {
	vals := make(map[rune]int, len(s))
	for _, c := range s {
		vals[c]++
	}
	return vals
}
