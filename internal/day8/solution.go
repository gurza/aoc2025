package day8

import (
	"sort"
	"strconv"
	"strings"
)

type Input struct {
	p1 uint64
	p2 uint64
}

type point struct {
	x, y, z int
}

type edge struct {
	i, j int
	d2   int
}

func Parse(input []string) Input {
	var p1, p2 uint64

	var points []point
	for _, line := range input {
		pp := strings.Split(line, ",")
		x, _ := strconv.Atoi(pp[0])
		y, _ := strconv.Atoi(pp[1])
		z, _ := strconv.Atoi(pp[2])
		points = append(points, point{x, y, z})
	}
	n := len(points)

	edges := make([]edge, 0, n*(n-1)/2)
	for i := range n {
		for j := i + 1; j < n; j++ {
			dx := points[i].x - points[j].x
			dy := points[i].y - points[j].y
			dz := points[i].z - points[j].z
			d2 := dx*dx + dy*dy + dz*dz
			edges = append(edges, edge{i, j, d2})
		}
	}

	sort.Slice(edges, func(a, b int) bool {
		if edges[a].d2 == edges[b].d2 {
			if edges[a].i == edges[b].i {
				return edges[a].j < edges[b].j
			}
			return edges[a].i < edges[b].i
		}
		return edges[a].d2 < edges[b].d2
	})

	return Input{p1, p2}
}

func Part1(input Input) uint64 {
	return input.p1
}

func Part2(input Input) uint64 {
	return input.p2
}
