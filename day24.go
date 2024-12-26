package main

import (
	"fmt"
	"maps"
	"slices"
	"strconv"
	"strings"
)

func Day24(input []string) {
	inputIndex := slices.Index(input, "")
	registers := map[string]int{}
	for _, register := range input[:inputIndex] {
		tmp := strings.Split(register, ": ")
		registers[tmp[0]], _ = strconv.Atoi(tmp[1])
	}
	type gate struct {
		a, b string
		op   string
	}
	inputPrefix := func(g gate, prefix string) bool {
		return strings.HasPrefix(g.a, prefix) || strings.HasPrefix(g.b, prefix)
	}
	gates := map[string]gate{}
	for _, calculation := range input[inputIndex+1:] {
		tmp := strings.Fields(calculation)
		gates[tmp[4]] = gate{tmp[0], tmp[2], tmp[1]}
	}
	var loadValue func(string) int
	loadValue = func(reg string) int {
		if val, ok := registers[reg]; ok {
			return val
		}
		gate := gates[reg]
		lh, rh := loadValue(gate.a), loadValue(gate.b)
		switch gate.op {
		case "AND":
			return lh & rh
		case "XOR":
			return lh ^ rh
		case "OR":
			return lh | rh
		default:
			panic("Unknown operand: " + gate.op)
		}
	}
	part1 := 0
	for reg := range maps.Keys(gates) {
		if reg[0] == 'z' {
			tmp, _ := strconv.Atoi(reg[1:])
			part1 |= loadValue(reg) << tmp
		}
	}
	fmt.Println(part1)
	bad := map[string]bool{}
	for res, g := range gates {
		if g.op != "XOR" && strings.HasPrefix(res, "z") && res != "z45" {
			bad[res] = true
		}
		if g.op == "XOR" &&
			!strings.HasPrefix(res, "z") &&
			!inputPrefix(g, "x") &&
			!inputPrefix(g, "y") {
			bad[res] = true
		}
		if g.op == "AND" && !inputPrefix(g, "x00") {
			for _, subg := range gates {
				if subg.op != "OR" && inputPrefix(subg, res) {
					bad[res] = true
				}
			}
		}
		if g.op == "XOR" {
			for _, subg := range gates {
				if subg.op == "OR" && inputPrefix(subg, res) {
					bad[res] = true
				}
			}
		}
	}
	fmt.Println(strings.Join(slices.Sorted(maps.Keys(bad)), ","))
}
