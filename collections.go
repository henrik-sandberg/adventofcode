package main

// Returns the index of e in s. -1 if s does not contain e
func IndexOf[T comparable](s []T, e T) int {
	for i, v := range s {
		if v == e {
			return i
		}
	}
	return -1
}

// Returns true if s contains e
func Contains[T comparable](s []T, e T) bool {
	return IndexOf(s, e) != -1
}

// Returns the count of e in s
func Count[T comparable](s []T, e T) (res int) {
	for _, v := range s {
		if v == e {
			res++
		}
	}
	return
}

// Returns a new slice after applying the predicate function
func Filter[T comparable](elems []T, predicate func(T) bool) []T {
	ret := make([]T, 0, len(elems))
	for _, e := range elems {
		if predicate(e) {
			ret = append(ret, e)
		}
	}
	return ret
}

// Returns the intersect of first and second as a new slice
func Intersect[T comparable](first, second []T) []T {
	a := toMap(first)
	b := toMap(second)
	ret := make([]T, 0, min(len(a), len(b)))
	for k, _ := range a {
		if _, ok := b[k]; ok {
			ret = append(ret, k)
		}
	}
	return ret
}

func toMap[T comparable](slice []T) map[T]bool {
	ret := make(map[T]bool, len(slice))
	for _, v := range slice {
		ret[v] = true
	}
	return ret
}

func bag(s string) map[rune]int {
	vals := make(map[rune]int, len(s))
        for _, c := range s {
		vals[c]++
	}
	return vals
}
