package y2023

import (
	"adventofcode/shared"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func Day07(input []string) (solution shared.Solution[int, int]) {
	sort.Slice(input, func(i, j int) bool {
		return isLess(
			strings.Fields(input[i])[0],
			strings.Fields(input[j])[0],
			plainOccurences,
			[]rune("23456789TJQKA"),
		)
	})
	solution.Part1 = sumCardBid(input)
	sort.Slice(input, func(i, j int) bool {
		return isLess(
			strings.Fields(input[i])[0],
			strings.Fields(input[j])[0],
			withJokers,
			[]rune("J23456789TQKA"),
		)
	})
	solution.Part2 = sumCardBid(input)
	return
}

func sumCardBid(sortedCards []string) (res int) {
	for i, s := range sortedCards {
		bid, _ := strconv.Atoi(strings.Fields(s)[1])
		res += (i + 1) * bid
	}
	return
}

func isLess(card1, card2 string, maxOccurrences func(string) []int, cardValue []rune) bool {
	a := maxOccurrences(card1)
	b := maxOccurrences(card2)
	for i := 0; i < min(len(a), len(b)); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	for i := 0; i < len(card1); i++ {
		ai := slices.Index(cardValue, rune(card1[i]))
		bi := slices.Index(cardValue, rune(card2[i]))
		if ai != bi {
			return ai < bi
		}
	}
	panic(fmt.Sprintf("Could not rank cards %v and %v", card1, card2))
}

func plainOccurences(s string) []int {
	vals := shared.Counts([]rune(s))
	counts := make([]int, 0, len(vals))
	for _, v := range vals {
		counts = append(counts, v)
	}
	sort.Slice(counts, func(i, j int) bool {
		return counts[i] > counts[j]
	})
	return counts
}

func withJokers(s string) []int {
	vals := shared.Counts([]rune(s))
	counts := make([]int, 0, len(vals))
	for k, v := range vals {
		if k != 'J' {
			counts = append(counts, v)
		}
	}
	sort.Slice(counts, func(i, j int) bool {
		return counts[i] > counts[j]
	})
	// Special case for only jokers
	if len(counts) == 0 {
		counts = append(counts, vals['J'])
	} else {
		counts[0] = min(5, counts[0]+vals['J'])
	}
	return counts
}
