package day6

import (
	"strconv"
	"strings"
)

type Input struct {
	p1 uint64
	p2 uint64
}

func Parse(input []string) Input {
	var p1, p2 uint64

	f0 := strings.Fields(input[0])
	cols1 := len(f0)
	rows := len(input)

	ff := make([][]string, rows)
	ff[0] = f0
	for r := 1; r < rows; r++ {
		ff[r] = strings.Fields(input[r])
	}

	// p1
	for c := range cols1 {
		nums := []uint64{}
		for r := 0; r < rows-1; r++ {
			num, _ := strconv.Atoi(ff[r][c])
			nums = append(nums, uint64(num))
		}

		sym := ff[rows-1][c][0]

		res := nums[0]
		if sym == '*' {
			for _, x := range nums[1:] {
				res *= x
			}
		} else {
			for _, x := range nums[1:] {
				res += x
			}
		}

		p1 += res
	}

	return Input{p1, p2}
}

func Part1(input Input) uint64 {
	return input.p1
}

func Part2(input Input) uint64 {
	return input.p2
}
