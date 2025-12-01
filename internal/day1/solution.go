package day1

import (
	"strconv"
)

func Parse(input []string) []int {
	res := make([]int, len(input))
	for i, s := range input {
		n, _ := strconv.Atoi(s[1:])
		if s[0] == 'L' {
			res[i] = -n
		} else {
			res[i] = n
		}
	}
	return res
}

func Part1(steps []int) int {
	pos := 50
	pwd := 0

	for _, step := range steps {
		pos += step
		if pos%100 == 0 {
			pwd += 1
		}
	}

	return pwd
}

func Part2(steps []int) int {
	pos := 50
	pwd := 0

	for _, step := range steps {
		old := pos
		pos = pos + step

		if step > 0 {
			// check (old, pos]
			pwd += floorDiv(pos, 100) - floorDiv(old, 100)
		} else {
			// check [old, pos)
			pwd += floorDiv(old-1, 100) - floorDiv(pos-1, 100)
		}
	}

	return pwd
}

// floorDiv returns floor(a/n) for positive n. Unlike Go's / operator,
// it rounds toward negative infinity (not toward zero).
func floorDiv(a, n int) int {
	if a >= 0 {
		return a / n
	}
	return (a - n + 1) / n
}
