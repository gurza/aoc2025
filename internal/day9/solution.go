package day9

import (
	"strconv"
	"strings"
)

type Input struct {
	p1 uint64
	p2 uint64
}

type point struct {
	x, y int
}

func Parse(input []string) Input {
	var p1, p2 uint64

	pts := make([]point, 0, len(input))
	for _, s := range input {
		pp := strings.Split(s, ",")
		x, _ := strconv.Atoi(pp[0])
		y, _ := strconv.Atoi(pp[1])
		pts = append(pts, point{x, y})
	}

	n := len(pts)
	for i := range n {
		for j := i + 1; j < n; j++ {
			dx := pts[i].x - pts[j].x
			if dx < 0 {
				dx = -dx
			}
			dx += 1
			dy := pts[i].y - pts[j].y
			if dy < 0 {
				dy = -dy
			}
			dy += 1

			a := uint64(dx) * uint64(dy)
			if a > p1 {
				p1 = a
			}
		}
	}

	return Input{p1, p2}
}

func Part1(input Input) uint64 {
	return input.p1
}

func Part2(input Input) uint64 {
	return input.p2
}
