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

// dsu implements a Disjoint Set Union structure.
type dsu struct {
	parent []int // parent pointer
	size   []int // size of the component represented by root
}

// newDSU initializes a DSU of n independent elements (0..n-1),
// each starting as its own component of size 1.
func newDSU(n int) *dsu {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range n {
		parent[i] = i
		size[i] = 1
	}
	return &dsu{parent, size}
}

// find returns the root of x's component.
func (d *dsu) find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.find(d.parent[x])
	}
	return d.parent[x]
}

// union merges the components containing a and b,
// attaching the smaller component under the larger one to keep trees shallow.
func (d *dsu) union(a, b int) {
	ra := d.find(a)
	rb := d.find(b)
	if ra == rb {
		return
	}
	if d.size[ra] < d.size[rb] {
		ra, rb = rb, ra
	}
	d.parent[rb] = ra
	d.size[ra] += d.size[rb]
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

	dsu := newDSU(n)
	limit := 1000 // todo: parametrize
	for i := range limit {
		e := edges[i]
		dsu.union(e.i, e.j)
	}

	seen := make(map[int]bool)
	var comps []int
	for i := range n {
		root := dsu.find(i)
		if !seen[root] {
			seen[root] = true
			comps = append(comps, dsu.size[root])
		}
	}
	sort.Slice(comps, func(i, j int) bool {
		return comps[i] > comps[j]
	})
	p1 = 1
	for i := range 3 {
		p1 *= uint64(comps[i])
	}

	return Input{p1, p2}
}

func Part1(input Input) uint64 {
	return input.p1
}

func Part2(input Input) uint64 {
	return input.p2
}
