package day2

import (
	"strconv"
	"strings"
)

type Input struct {
	p1 uint64
	p2 uint64
}

func Parse(input []string) Input {
	var sum1, sum2 uint64

	for r := range strings.SplitSeq(input[0], ",") {
		pp := strings.Split(r, "-")
		from, _ := strconv.Atoi(pp[0])
		to, _ := strconv.Atoi(pp[1])

		for num := from; num <= to; num++ {
			s := strconv.Itoa(num)
			len := len(s)

			// repeated twice
			half := len / 2
			if s[:half] == s[half:] {
				sum1 += uint64(num)
				sum2 += uint64(num)
				continue
			}

			// repeated
			for lenU := 1; lenU <= len/2; lenU++ {
				if len%lenU != 0 {
					continue
				}
				unit := s[:lenU]
				repeated := strings.Repeat(unit, len/lenU)
				if s == repeated {
					sum2 += uint64(num)
					break
				}
			}
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
