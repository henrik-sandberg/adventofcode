package main

import (
	"fmt"
	"regexp"
	"strings"
)

func Day04(input []string) {
	m := Map{strings.Join(input, ""), len(input[0])}
	fmt.Println(day04_part1(m))
	fmt.Println(day04_part2(m))
}

func day04_part1(m Map) (result int) {
	search := "XMAS"
	searchReverse := Reverse(search)
	max_row := len(m.cells) / m.width
	max_col := m.width
	rows := make([]strings.Builder, max_row)
	cols := make([]strings.Builder, max_col)
	fdiag := make([]strings.Builder, max_row+max_col-1)
	bdiag := make([]strings.Builder, len(fdiag))
	for i := range m.cells {
		rows[i/m.width].WriteByte(m.cells[i])
		cols[i%m.width].WriteByte(m.cells[i])
		fdiag[i%m.width+i/m.width].WriteByte(m.cells[i])
		bdiag[i%m.width-i/m.width+max_row-1].WriteByte(m.cells[i])
	}
	for _, arr := range [][]strings.Builder{rows, cols, fdiag, bdiag} {
		for _, sb := range arr {
			result += strings.Count(sb.String(), search) + strings.Count(sb.String(), searchReverse)
		}
	}
	return
}

func day04_part2(m Map) (result int) {
	re := regexp.MustCompile("A")
	for _, ind := range re.FindAllStringIndex(m.cells, -1) {
		if hasMasNeighbours(m, ind[0]) {
			result++
		}
	}
	return
}

func hasMasNeighbours(m Map, ind int) bool {
	if ind/m.width == 0 ||
		ind/m.width == len(m.cells)/m.width-1 ||
		ind%m.width == 0 ||
		ind%m.width == m.width-1 {
		return false
	}
	ns := []byte{
		m.cells[ind-m.width-1],
		m.cells[ind-m.width+1],
		m.cells[ind+m.width-1],
		m.cells[ind+m.width+1],
	}
	return ns[0] != ns[3] && Count(ns, 'M') == 2 && Count(ns, 'S') == 2
}
