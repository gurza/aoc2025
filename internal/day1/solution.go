package day1

import (
	"strconv"
)

type Input struct {
	p1 uint64
	p2 uint64
}

func Parse(input []string) Input {
	var sum1, sum2 uint64

	pos := 50
	old := 0
	for _, s := range input {
		num, _ := strconv.Atoi(s[1:])

		old = pos
		if s[0] == 'L' {
			pos -= num
			// check [old, pos)
			sum2 += uint64(floorDiv(old-1, 100) - floorDiv(pos-1, 100))
		} else {
			pos += num
			// check (old, pos]
			sum2 += uint64(floorDiv(pos, 100) - floorDiv(old, 100))
		}

		if pos%100 == 0 {
			sum1 += 1
		}
	}

	return Input{sum1, sum2}
}

func Part1(input Input) uint64 {
	return input.p1
}

func Part2(input Input) uint64 {
	return input.p2
}

// floorDiv returns floor(a/n) for positive n. Unlike Go's / operator,
// it rounds toward negative infinity (not toward zero).
func floorDiv(a, n int) int {
	if a >= 0 {
		return a / n
	}
	return (a - n + 1) / n
}
