package y2022

import (
	"adventofcode/solutions/shared"
	"encoding/json"
	"fmt"
	"reflect"
	"slices"
	"sort"
)

func Day13(input []string) (solution shared.Solution[int, int]) {
	var compare func(left []any, right []any) int
	compare = func(left []any, right []any) int {
		for i := 0; i < len(left) && i < len(right); i++ {
			l := left[i]
			r := right[i]
			typeLeft := reflect.TypeOf(l).Kind()
			typeRight := reflect.TypeOf(r).Kind()
			if typeLeft == typeRight {
				if typeLeft == reflect.Float64 {
					if res := int(l.(float64) - r.(float64)); res != 0 {
						return res
					}
				} else if res := compare(l.([]any), r.([]any)); res != 0 {
					return res
				}
			} else {
				var res int
				if typeLeft == reflect.Float64 {
					res = compare([]any{l.(float64)}, r.([]any))
				} else {
					res = compare(l.([]any), []any{r.(float64)})
				}
				if res != 0 {
					return res
				}
			}
		}
		return len(left) - len(right)
	}
	for i := 0; i < len(input); i = i + 3 {
		var left []any
		var right []any
		json.Unmarshal([]byte(input[i]), &left)
		json.Unmarshal([]byte(input[i+1]), &right)
		if compare(left, right) < 0 {
			index := i/3 + 1
			solution.Part1 += index
		}
	}
	var packets [][]any
	for _, line := range input {
		if line != "" {
			var packet []any
			json.Unmarshal([]byte(line), &packet)
			packets = append(packets, packet)
		}
	}
	dividerPackets := [][]any{{[]any{float64(2)}}, {[]any{float64(6)}}}
	packets = append(packets, dividerPackets...)

	sort.Slice(packets, func(i, j int) bool {
		return compare(packets[i], packets[j]) < 0
	})

	strings := make([]string, len(packets))
	for i, p := range packets {
		strings[i] = fmt.Sprint(p)
	}

	indexes := make([]int, len(dividerPackets))
	for i, dp := range dividerPackets {
		indexes[i] = slices.Index(strings, fmt.Sprint(dp)) + 1
	}
	solution.Part2 = shared.Product(indexes...)
	return
}
