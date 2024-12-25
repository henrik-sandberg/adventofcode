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
	gates := map[string][]string{}
	for _, calculation := range input[inputIndex+1:] {
		tmp := strings.Fields(calculation)
		gates[tmp[4]] = tmp[:3]
	}
	operands := map[string]func(int, int) int{
		"AND": func(a, b int) int {
			return a & b
		},
		"XOR": func(a, b int) int {
			return a ^ b
		},
		"OR": func(a, b int) int {
			return a | b
		},
	}
	var loadValue func(string) int
	loadValue = func(reg string) int {
		if v, ok := registers[reg]; ok {
			return v
		}
		gate := gates[reg]
		v := operands[gate[1]](loadValue(gate[0]), loadValue(gate[2]))
		registers[reg] = v
		return v
	}
	Z := []string{}
	for reg := range maps.Keys(gates) {
		if reg[0] == 'z' {
			Z = append(Z, reg)
		}
	}
	slices.Sort(Z)
	part1 := 0
	for _, reg := range Z {
		tmp, _ := strconv.Atoi(reg[1:])
		part1 |= loadValue(reg) << tmp
	}
	fmt.Println(part1)
	sumRegistry := func(prefix byte) (res int) {
		for reg := range maps.Keys(registers) {
			if reg[0] == prefix {
				tmp, _ := strconv.Atoi(reg[1:])
				res |= loadValue(reg) << tmp
			}
		}
		return
	}
	fmt.Println(sumRegistry('x'))

	/*
		sumRegistry := func(prefix byte) (res int) {
			for reg := range maps.Keys(registers) {
				if reg[0] == prefix {
					tmp, _ := strconv.Atoi(reg[1:])
					res |= loadValue(reg) << tmp
				}
			}
			return
		}
		expected := sumRegistry('x') & sumRegistry('y')
		fmt.Println("Expected:", expected)
		zRegisters := []string{}
		for reg := range maps.Keys(gates) {
			if reg[0] == 'z' {
				zRegisters = append(zRegisters, reg)
			}
		}
		keys := slices.Collect(maps.Keys(gates))
		fmt.Println("z registers:", zRegisters)
		for perm := range Permutations(keys, 8) {
			registers = map[string]int{}
			for _, register := range input[:inputIndex] {
				tmp := strings.Split(register, ": ")
				registers[tmp[0]], _ = strconv.Atoi(tmp[1])
			}
			gates[perm[0]], gates[perm[1]] = gates[perm[1]], gates[perm[0]]
			gates[perm[2]], gates[perm[3]] = gates[perm[3]], gates[perm[2]]
			gates[perm[4]], gates[perm[5]] = gates[perm[5]], gates[perm[4]]
			gates[perm[6]], gates[perm[7]] = gates[perm[7]], gates[perm[6]]
			res := 0
			for _, reg := range zRegisters {
				tmp, _ := strconv.Atoi(reg[1:])
				res |= loadValue(reg) << tmp
			}
			if res == expected {
				fmt.Println(perm, res)
				panic("found")
			}
			gates[perm[0]], gates[perm[1]] = gates[perm[1]], gates[perm[0]]
			gates[perm[2]], gates[perm[3]] = gates[perm[3]], gates[perm[2]]
			gates[perm[4]], gates[perm[5]] = gates[perm[5]], gates[perm[4]]
			gates[perm[6]], gates[perm[7]] = gates[perm[7]], gates[perm[6]]
		}
	*/
}
