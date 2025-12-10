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

type tile struct {
	x, y int
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

func Part1(input Input) uint64 {
	return input.p1
}

func Part2(input Input) uint64 {
	return input.p2
}
