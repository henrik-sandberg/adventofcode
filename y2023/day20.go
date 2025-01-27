package y2023

import (
	"adventofcode/shared"
	"slices"
	"strings"
)

type module interface {
	send(bool, string)
}

type button struct {
	queue   *[]func()
	pulses  *map[bool]int
	modules *map[string]module
}

type broadcaster struct {
	queue   *[]func()
	modules *map[string]module
	pulses  *map[bool]int
	id      string
	output  []string
}

type flipflop struct {
	queue   *[]func()
	modules *map[string]module
	pulses  *map[bool]int
	id      string
	output  []string
	on      bool
}

type conjunction struct {
	queue   *[]func()
	modules *map[string]module
	pulses  *map[bool]int
	id      string
	output  []string
	input   map[string]bool
}

func (b *button) press() {
	//fmt.Println("id: button sending pulse false to broadcaster")
	(*b.pulses)[false]++
	(*b.queue) = append(*b.queue, func() { (*b.modules)["broadcaster"].send(false, "button") })
}

func (b *broadcaster) send(pulse bool, id string) {
	for _, out := range b.output {
		//fmt.Printf("id: %s sending pulse %v to %v\n", b.id, pulse, out)
		(*b.pulses)[pulse]++
		(*b.queue) = append(*b.queue, func() { (*b.modules)[out].send(pulse, b.id) })
	}
}

func (f *flipflop) send(pulse bool, id string) {
	if !pulse {
		f.on = !f.on
		for _, out := range f.output {
			//fmt.Printf("id: %s sending pulse %v to %v\n", f.id, f.on, out)
			(*f.pulses)[f.on]++
			if m, ok := (*f.modules)[out]; ok {
				(*f.queue) = append(*f.queue, func() { m.send(f.on, f.id) })
			}
		}
	}
}

func (c *conjunction) send(pulse bool, id string) {
	c.input[id] = pulse
	all := true
	for _, v := range c.input {
		all = all && v
	}
	for _, out := range c.output {
		(*c.pulses)[!all]++
		//fmt.Printf("id: %s sending pulse %v to %v\n", c.id, !all, out)
		if m, ok := (*c.modules)[out]; ok {
			(*c.queue) = append(*c.queue, func() { m.send(!all, c.id) })
		}
	}
}

func Day20(input []string) (solution shared.Solution[int, int]) {
	queue := []func(){}
	modules := map[string]module{}
	pulses := map[bool]int{}
	for _, line := range input {
		f := strings.Split(line, " -> ")
		targets := strings.Split(f[1], ", ")
		if f[0] == "broadcaster" {
			modules[f[0]] = &broadcaster{queue: &queue, modules: &modules, pulses: &pulses, id: f[0], output: targets}
		} else if f[0][0] == '%' {
			id := f[0][1:]
			modules[id] = &flipflop{queue: &queue, modules: &modules, pulses: &pulses, id: id, output: targets}
		} else if f[0][0] == '&' {
			id := f[0][1:]
			modules[id] = &conjunction{queue: &queue, modules: &modules, pulses: &pulses, id: id, output: targets, input: map[string]bool{}}
		}
	}
	for _, m := range modules {
		if con, ok := m.(*conjunction); ok {
			for _, m := range modules {
				f, ok := m.(*flipflop)
				if ok && slices.Contains(f.output, con.id) {
					con.input[f.id] = false
				}
			}
		}
	}
	button := button{queue: &queue, modules: &modules, pulses: &pulses}
	for i := 0; i < 1e3; i++ {
		button.press()
		for len(queue) > 0 {
			queue[0]()
			queue = queue[1:]
		}
	}
	solution.Part1 = pulses[true] * pulses[false]
	return
}
