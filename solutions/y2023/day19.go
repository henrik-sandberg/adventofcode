package y2023

import (
	"adventofcode/solutions/shared"
	"regexp"
	"strconv"
	"strings"
)

func Day19(input []string) (solution shared.Solution[int, int]) {
	type Part map[string]int
	type Workflow func(Part) string
	workflows := map[string][]Workflow{}
	ind := 0
	for ; input[ind] != ""; ind++ {
		tmp := strings.Split(input[ind], "{")
		name := tmp[0]
		for _, rule := range strings.Split(tmp[1][:len(tmp[1])-1], ",") {
			switch tmp := strings.Split(rule, ":"); len(tmp) {
			case 1:
				workflows[name] = append(workflows[name], func(_ Part) string {
					return tmp[0]
				})
			case 2:
				cond := tmp[0]
				c := cond[:1]
				op := cond[1]
				val, _ := strconv.Atoi(cond[2:])
				workflows[name] = append(workflows[name], func(p Part) string {
					if op == '<' && p[c] < val || op == '>' && p[c] > val {
						return tmp[1]
					}
					return ""
				})
			}
		}
	}
	var process func(Part, []Workflow) string
	process = func(p Part, wfs []Workflow) string {
		for _, wf := range wfs {
			switch val := wf(p); val {
			case "":
				continue
			case "A", "R":
				return val
			default:
				return process(p, workflows[val])
			}
		}
		return ""
	}
	numsRegexp := regexp.MustCompile(`\d+`)
	for i := ind + 1; i < len(input); i++ {
		nums := shared.IntSlice(numsRegexp.FindAllString(input[i], -1))
		part := Part{
			"x": nums[0],
			"m": nums[1],
			"a": nums[2],
			"s": nums[3],
		}
		if process(part, workflows["in"]) == "A" {
			solution.Part1 += shared.Sum(nums...)
		}
	}
	return
}
