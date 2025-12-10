package day10

import (
	"strconv"
	"strings"
)

type Input struct {
	p1 uint64
	p2 uint64
}

type machine struct {
	lights   int
	buttons  []int
	joltages []int
}

func Parse(input []string) Input {
	var p1, p2 uint64

	for _, line := range input {
		parse(line)
	}

	return Input{p1, p2}
}

func parse(s string) machine {
	fs := strings.Fields(s)

	// lights (fs[0])
	var ls int
	bs := []byte(fs[0][1:]) // skip '['
	for i, b := range bs {
		if b == '#' {
			ls |= 1 << i
		}
	}

	last := len(fs) - 1

	// buttons (fs[1..n-1])
	var btns []int
	for _, t := range fs[1:last] {
		t = t[1 : len(t)-1]
		var mask int
		for p := range strings.SplitSeq(t, ",") {
			i, _ := strconv.Atoi(p)
			mask |= 1 << i
		}
		btns = append(btns, mask)
	}

	// joltages {fs[last]}
	var js []int
	jl := fs[last]
	jl = jl[1 : len(jl)-1]
	for p := range strings.SplitSeq(jl, ",") {
		j, _ := strconv.Atoi(p)
		js = append(js, j)
	}

	return machine{ls, btns, js}
}

func Part1(input Input) uint64 {
	return input.p1
}

func Part2(input Input) uint64 {
	return input.p2
}
