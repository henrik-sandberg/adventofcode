package y2023

import (
	"adventofcode/solutions/shared"
	"regexp"
	"strconv"
	"strings"
)

func Day19(input []string) (solution shared.Solution[int, int]) {
	type Range struct {
		Lo, Hi int
	}
	type Rule struct {
		Category byte
		Op       byte
		Value    int
		Target   string
	}
	workflows := make(map[string][]Rule)
	idx := 0
	for ; input[idx] != ""; idx++ {
		tmp := strings.Split(input[idx], "{")
		name := tmp[0]
		var wf []Rule
		for _, rule := range strings.Split(tmp[1][:len(tmp[1])-1], ",") {
			parts := strings.Split(rule, ":")
			if len(parts) == 1 {
				wf = append(wf, Rule{
					Target: parts[0],
				})
			} else {
				cond := parts[0]
				val, _ := strconv.Atoi(cond[2:])
				wf = append(wf, Rule{
					Category: cond[0],
					Op:       cond[1],
					Value:    val,
					Target:   parts[1],
				})
			}
		}
		workflows[name] = wf
	}

	splitRange := func(pr map[byte]Range, rule Rule) (map[byte]Range, map[byte]Range) {
		r := pr[rule.Category]
		var tLo, tHi, fLo, fHi int
		switch rule.Op {
		case '<':
			tLo, tHi = r.Lo, min(r.Hi, rule.Value-1)
			fLo, fHi = max(r.Lo, rule.Value), r.Hi
		case '>':
			tLo, tHi = max(r.Lo, rule.Value+1), r.Hi
			fLo, fHi = r.Lo, min(r.Hi, rule.Value)
		}
		var truePr, falsePr map[byte]Range
		if tLo <= tHi {
			truePr = make(map[byte]Range, len(pr))
			for k, v := range pr {
				truePr[k] = v
			}
			truePr[rule.Category] = Range{tLo, tHi}
		}
		if fLo <= fHi {
			falsePr = make(map[byte]Range, len(pr))
			for k, v := range pr {
				falsePr[k] = v
			}
			falsePr[rule.Category] = Range{fLo, fHi}
		}
		return truePr, falsePr
	}

	var dfs func(string, map[byte]Range) int
	dfs = func(wfName string, pr map[byte]Range) int {
		if wfName == "A" {
			return (pr['x'].Hi - pr['x'].Lo + 1) *
				(pr['m'].Hi - pr['m'].Lo + 1) *
				(pr['a'].Hi - pr['a'].Lo + 1) *
				(pr['s'].Hi - pr['s'].Lo + 1)
		}
		if wfName == "R" {
			return 0
		}
		total := 0
		for _, rule := range workflows[wfName] {
			truePr, falsePr := splitRange(pr, rule)
			if len(truePr) > 0 {
				total += dfs(rule.Target, truePr)
			}
			if len(falsePr) == 0 {
				return total
			}
			pr = falsePr
		}
		return total
	}
	numsRegexp := regexp.MustCompile(`\d+`)
	idx++
	for ; idx < len(input); idx++ {
		nums := shared.IntSlice(numsRegexp.FindAllString(input[idx], 4))
		part := map[byte]Range{
			'x': {Lo: nums[0], Hi: nums[0]},
			'm': {Lo: nums[1], Hi: nums[1]},
			'a': {Lo: nums[2], Hi: nums[2]},
			's': {Lo: nums[3], Hi: nums[3]},
		}
		if dfs("in", part) > 0 {
			solution.Part1 += part['x'].Lo + part['m'].Lo + part['a'].Lo + part['s'].Lo
		}
	}
	solution.Part2 = dfs("in", map[byte]Range{
		'x': {Lo: 1, Hi: 4000},
		'm': {Lo: 1, Hi: 4000},
		'a': {Lo: 1, Hi: 4000},
		's': {Lo: 1, Hi: 4000},
	})
	return
}
