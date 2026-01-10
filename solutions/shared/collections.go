package shared

import "iter"

// Returns the count of elements in s as a map
func Counts[T comparable](s []T) map[T]int {
	m := make(map[T]int)
	for _, v := range s {
		m[v]++
	}
	return m
}

// Returns the count of e in s
func Count[T comparable](s []T, e T) int {
	res := 0
	for _, v := range s {
		if v == e {
			res++
		}
	}
	return res
}

// Returns the intersect of first and second as a new slice
func Intersect[T comparable](first, second []T) []T {
	a := toMap(first)
	b := toMap(second)
	ret := make([]T, 0, min(len(a), len(b)))
	for k := range a {
		if _, ok := b[k]; ok {
			ret = append(ret, k)
		}
	}
	return ret
}

func Combinations[T any](set []T, length int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		work := make([]T, length)
		var helper func(int, int) bool
		helper = func(start, depth int) bool {
			if depth == length {
				out := make([]T, length)
				copy(out, work)
				return yield(out)
			}
			for i := start; i <= len(set)-(length-depth); i++ {
				work[depth] = set[i]
				if !helper(i+1, depth+1) {
					return false
				}
			}
			return true
		}
		helper(0, 0)
	}
}

func Permutations[T any](set []T, length int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		used := make([]bool, len(set))
		work := make([]T, length)
		var helper func(int) bool
		helper = func(k int) bool {
			if k == length {
				out := make([]T, length)
				copy(out, work)
				return yield(out)
			}
			for i := range set {
				if !used[i] {
					used[i] = true
					work[k] = set[i]
					if !helper(k + 1) {
						return false
					}
					used[i] = false
				}
			}
			return true
		}
		helper(0)
	}
}

// Returns an iterator over the idx indexed values of sub-slice
func Column[T any](sl [][]T, idx int) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, s := range sl {
			if !yield(s[idx]) {
				return
			}
		}
	}
}

func toMap[T comparable](slice []T) map[T]bool {
	ret := make(map[T]bool, len(slice))
	for _, v := range slice {
		ret[v] = true
	}
	return ret
}
