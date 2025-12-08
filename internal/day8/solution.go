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
// It returns true if a merge actually happened.
func (d *dsu) union(a, b int) bool {
	ra := d.find(a)
	rb := d.find(b)
	if ra == rb {
		return false
	}
	if d.size[ra] < d.size[rb] {
		ra, rb = rb, ra
	}
	d.parent[rb] = ra
	d.size[ra] += d.size[rb]
	return true
}

func Parse(input []string) Input {
	var p1, p2 uint64

	var pts []point
	for _, s := range input {
		pp := strings.Split(s, ",")
		x, _ := strconv.Atoi(pp[0])
		y, _ := strconv.Atoi(pp[1])
		z, _ := strconv.Atoi(pp[2])
		pts = append(pts, point{x, y, z})
	}
	n := len(pts)

	es := make([]edge, 0, n*(n-1)/2)
	for i := range n {
		for j := i + 1; j < n; j++ {
			dx := pts[i].x - pts[j].x
			dy := pts[i].y - pts[j].y
			dz := pts[i].z - pts[j].z
			d2 := dx*dx + dy*dy + dz*dz
			es = append(es, edge{i, j, d2})
		}
	}

	sort.Slice(es, func(a, b int) bool {
		if es[a].d2 == es[b].d2 {
			if es[a].i == es[b].i {
				return es[a].j < es[b].j
			}
			return es[a].i < es[b].i
		}
		return es[a].d2 < es[b].d2
	})

	dsu := newDSU(n)
	lim := 1000
	if len(es) < lim { // FIXME: parametrize limit
		lim = n / 2
	}

	comp := n
	k := 0
	done1 := false

	var last edge
	done2 := false

	for _, e := range es {
		k++
		if dsu.union(e.i, e.j) {
			comp--
			if !done2 && comp == 1 {
				last = e
				done2 = true
			}
		}

		if !done1 && k == lim {
			p1 = calcP1(dsu, n)
			done1 = true
		}

		if done1 && done2 {
			break
		}
	}

	if !done1 {
		p1 = calcP1(dsu, n)
	}
	if done2 {
		p2 = uint64(pts[last.i].x * pts[last.j].x)
	}

	return Input{p1, p2}
}

func calcP1(d *dsu, n int) uint64 {
	seen := make(map[int]bool, n)
	comps := make([]int, 0, n)

	for i := 0; i < n; i++ {
		r := d.find(i)
		if !seen[r] {
			seen[r] = true
			comps = append(comps, d.size[r])
		}
	}

	sort.Slice(comps, func(i, j int) bool {
		return comps[i] > comps[j]
	})

	res := uint64(1)
	for i := 0; i < 3 && i < len(comps); i++ {
		res *= uint64(comps[i])
	}
	return res
}

func Part1(input Input) uint64 {
	return input.p1
}

func Part2(input Input) uint64 {
	return input.p2
}
