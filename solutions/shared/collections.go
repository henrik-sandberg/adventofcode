package shared

// Returns the count of elements in s as a map
func Counts[T comparable](s []T) map[T]int {
	m := map[T]int{}
	for _, v := range s {
		m[v]++
	}
	return m
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

func Combinations[T interface{}](set []T, length int) <-chan []T {
	ch := make(chan []T)

	go func() {
		defer close(ch)
		var helper func([]T, int, []T)

		helper = func(currentSet []T, length int, currentComb []T) {
			if length == 0 {
				comb := make([]T, len(currentComb))
				copy(comb, currentComb)
				ch <- comb
				return
			}
			for i := 0; i <= len(currentSet)-length; i++ {
				helper(currentSet[i+1:], length-1, append(currentComb, currentSet[i]))
			}
		}
		helper(set, length, []T{})
	}()
	return ch
}

func Permutations[T interface{}](set []T, length int) <-chan []T {
	ch := make(chan []T)
	go func() {
		defer close(ch)
		used := make([]bool, len(set))
		temp := make([]T, length)
		var helper func(int)
		helper = func(k int) {
			if k == length {
				tmp := make([]T, length)
				copy(tmp, temp)
				ch <- tmp
				return
			}
			for i := 0; i < len(set); i++ {
				if !used[i] {
					used[i] = true
					temp[k] = set[i]
					helper(k + 1)
					used[i] = false
				}
			}
		}
		helper(0)
	}()
	return ch
}

func toMap[T comparable](slice []T) map[T]bool {
	ret := make(map[T]bool, len(slice))
	for _, v := range slice {
		ret[v] = true
	}
	return ret
}
