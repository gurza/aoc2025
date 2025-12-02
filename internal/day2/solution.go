package day2

import (
	"strconv"
	"strings"
)

func Parse(input []string) uint64 {
	var sum uint64

	for r := range strings.SplitSeq(input[0], ",") {
		pp := strings.Split(r, "-")
		from, _ := strconv.Atoi(pp[0])
		to, _ := strconv.Atoi(pp[1])

		for num := from; num <= to; num++ {
			if isInvalid(num) {
				sum += uint64(num)
			}
		}
	}

	return sum
}

func isInvalid(n int) bool {
	s := strconv.Itoa(n)
	l := len(s)
	if l%2 != 0 {
		return false
	}
	half := l / 2
	return s[:half] == s[half:]
}

func Part1(p1 uint64) uint64 {
	return p1
}
