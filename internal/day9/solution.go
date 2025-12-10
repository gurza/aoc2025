package day9

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

type Input struct {
	p1 uint64
	p2 uint64
}

const (
	outside = 0
	inside  = 1
	unknown = 2
)

type tile struct {
	x, y int
}

// grid is a 2D array stored in a single backing slice (row-major).
type grid struct {
	width, height int
	values        []int
}

// newGrid creates a width x height grid filled with the given value.
func newGrid(width, height, value int) *grid {
	g := &grid{
		width:  width,
		height: height,
		values: make([]int, width*height),
	}
	for i := range g.values {
		g.values[i] = value
	}
	return g
}

// index computes the 1D index for a tile in row-major order.
func (g *grid) index(t tile) int {
	return t.y*g.width + t.x
}

// set writes the value at the given tile.
func (g *grid) set(t tile, v int) {
	g.values[g.index(t)] = v
}

func Parse(input []string) Input {
	var p1, p2 uint64

	tiles := make([]tile, 0, len(input))
	for _, s := range input {
		pp := strings.Split(s, ",")
		x, _ := strconv.Atoi(pp[0])
		y, _ := strconv.Atoi(pp[1])
		tiles = append(tiles, tile{x, y})
	}

	n := len(tiles)
	for i := range n {
		for j := i + 1; j < n; j++ {
			dx := absInt(tiles[i].x-tiles[j].x) + 1
			dy := absInt(tiles[i].y-tiles[j].y) + 1
			a := uint64(dx) * uint64(dy)
			if a > p1 {
				p1 = a
			}
		}
	}

	shrinkX := shrink(tiles, 0)
	shrinkY := shrink(tiles, 1)
	shrunk := make([]tile, len(tiles))
	for i, t := range tiles {
		shrunk[i] = tile{
			x: shrinkX[t.x],
			y: shrinkY[t.y],
		}
	}
	g := newGrid(len(shrinkX), len(shrinkY), unknown)
	rasterizePolygonEdges(shrunk, g)

	return Input{p1, p2}
}

func absInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// shrink performs 1D coordinate compression for either the x- or y-axis.
//
// It takes a slice of tiles and an axis selector:
//   - axis == 0 -> use tile.x
//   - axis == 1 -> use tile.y
func shrink(tiles []tile, axis int) map[int]int {
	coords := make([]int, 0, len(tiles)+2)
	for _, t := range tiles {
		if axis == 0 {
			coords = append(coords, t.x)
		} else {
			coords = append(coords, t.y)
		}
	}
	coords = append(coords, math.MinInt)
	coords = append(coords, math.MaxInt)

	sort.Ints(coords)

	// deduplication
	n := 0
	for i := 0; i < len(coords); i++ {
		if n == 0 || coords[i] != coords[n-1] {
			coords[n] = coords[i]
			n++
		}
	}
	coords = coords[:n]

	res := make(map[int]int, len(coords))
	for i, v := range coords {
		res[v] = i
	}
	return res
}

// edgeBounds returns the bounding box coordinates (x1,y1,x2,y2) of two tiles.
func edgeBounds(a, b tile) (x1, y1, x2, y2 int) {
	if a.x < b.x {
		x1, x2 = a.x, b.x
	} else {
		x1, x2 = b.x, a.x
	}
	if a.y < b.y {
		y1, y2 = a.y, b.y
	} else {
		y1, y2 = b.y, a.y
	}
	return
}

// rasterizePolygonEdges marks the polygon edges as `inside` on the
// compressed grid.
func rasterizePolygonEdges(shrunk []tile, g *grid) {
	size := len(shrunk)
	for i := range size {
		a := shrunk[i]
		b := shrunk[(i+1)%size]

		x1, y1, x2, y2 := edgeBounds(a, b)

		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				g.set(tile{x: x, y: y}, inside)
			}
		}
	}
}

func Part1(input Input) uint64 {
	return input.p1
}

func Part2(input Input) uint64 {
	return input.p2
}
