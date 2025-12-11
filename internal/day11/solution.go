package day11

import (
	"math"
	"strings"
)

type Input struct {
	p1 uint64
	p2 uint64
}

func Parse(input []string) Input {
	var p1, p2 uint64

	graph := make([][]uint, 26*26*26) // all three letters combinations
	for _, line := range input {
		pp := strings.Split(line, ":")
		src := pp[0]
		srcIdx := idx(src)
		for dst := range strings.FieldsSeq(pp[1]) {
			dstIdx := idx(dst)
			graph[srcIdx] = append(graph[srcIdx], dstIdx)
		}
	}

	p1 = uint64(paths(graph, "you", "out"))

	p2 = uint64(paths(graph, "svr", "fft") * paths(graph, "fft", "dac") * paths(graph, "dac", "out"))
	p2 += uint64(paths(graph, "svr", "dac") * paths(graph, "dac", "fft") * paths(graph, "fft", "out"))

	return Input{p1, p2}
}

func paths(graph [][]uint, start, end string) uint {
	cache := make([]uint, len(graph))
	for i := range cache {
		cache[i] = math.MaxUint
	}
	return dfs(graph, cache, idx(start), idx(end))
}

func dfs(graph [][]uint, cache []uint, node, end uint) uint {
	if node == end {
		return 1
	}
	if cache[node] != math.MaxUint {
		return cache[node]
	}

	var sum uint
	for _, next := range graph[node] {
		sum += dfs(graph, cache, next, end)
	}
	cache[node] = sum
	return sum
}

func idx(s string) uint {
	var res uint = 0
	for i := range 3 {
		res = 26*res + uint(s[i]-'a')
	}
	return res
}

func Part1(input Input) uint64 {
	return input.p1
}

func Part2(input Input) uint64 {
	return input.p2
}
