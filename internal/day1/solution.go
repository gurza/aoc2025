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
