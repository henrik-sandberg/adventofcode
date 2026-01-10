package y2022

import (
	"adventofcode/solutions/shared"
	"strconv"
	"strings"
)

func Day21(input []string) (solution shared.Solution[int, int]) {
	type Expr struct {
		left, right string
		op          byte
		val         int
		isValue     bool
	}
	exprs := make(map[string]Expr)
	for _, line := range input {
		id, s, _ := strings.Cut(line, ": ")
		n, err := strconv.Atoi(s)
		if err == nil {
			exprs[id] = Expr{val: n, isValue: true}
		} else {
			fs := strings.Fields(s)
			exprs[id] = Expr{left: fs[0], op: fs[1][0], right: fs[2]}
		}
	}
	values := make(map[string]int)
	resolvable := make(map[string]bool)
	var eval func(string) int
	eval = func(node string) int {
		e := exprs[node]
		if e.isValue {
			values[node] = e.val
			resolvable[node] = node != "humn"
			return e.val
		}
		left := eval(e.left)
		right := eval(e.right)
		resolvable[node] = resolvable[e.left] && resolvable[e.right]
		var v int
		switch e.op {
		case '+':
			v = left + right
		case '-':
			v = left - right
		case '*':
			v = left * right
		case '/':
			v = left / right
		}
		values[node] = v
		return v
	}
	solution.Part1 = eval("root")

	invertLeft := func(op byte, right, target int) int {
		switch op {
		case '+':
			return target - right
		case '-':
			return target + right
		case '*':
			return target / right
		case '/':
			return target * right
		}
		panic(op)
	}
	invertRight := func(op byte, left, target int) int {
		switch op {
		case '+':
			return target - left
		case '-':
			return left - target
		case '*':
			return target / left
		case '/':
			return left / target
		}
		panic(op)
	}

	var eval2 func(string, int) int
	eval2 = func(node string, target int) int {
		if node == "humn" {
			return target
		}
		e := exprs[node]
		if !resolvable[e.left] {
			return eval2(e.left, invertLeft(e.op, values[e.right], target))
		}
		return eval2(e.right, invertRight(e.op, values[e.left], target))
	}
	root := exprs["root"]
	if resolvable[root.left] {
		solution.Part2 = eval2(root.right, values[root.left])
	} else {
		solution.Part2 = eval2(root.left, values[root.right])
	}
	return
}
