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

		op := ff[rows-1][c][0]

		res := nums[0]
		if op == '*' {
			for _, x := range nums[1:] {
				res *= x
			}
		} else { // +
			for _, x := range nums[1:] {
				res += x
			}
		}

		p1 += res
	}

	cols2 := len(input[0]) // all strings in input have equal length

	grid := make([][]byte, rows)
	for r, line := range input {
		grid[r] = []byte(line)
	}

	solveP2 := func(colIdxs []int) {
		var op byte
		for _, c := range colIdxs {
			b := grid[rows-1][c]
			if b == '*' || b == '+' {
				op = b
				break
			}
		}

		nums := make([]uint64, 0, len(colIdxs))
		for _, c := range colIdxs {
			var v uint64
			hasDigit := false
			for r := 0; r < rows-1; r++ {
				ch := grid[r][c]
				if ch == ' ' {
					continue
				}
				hasDigit = true
				v = v*10 + uint64(ch-'0')
			}

			if hasDigit {
				nums = append(nums, v)
			}
		}

		res := nums[0]
		if op == '*' {
			for _, x := range nums[1:] {
				res *= x
			}
		} else { // '+'
			for _, x := range nums[1:] {
				res += x
			}
		}

		p2 += res
	}

	// p2
	cur := make([]int, 0) // current problem
	for c := cols2 - 1; c >= 0; c-- {
		sep := true
		for r := range rows {
			if grid[r][c] != ' ' {
				sep = false
				break
			}
		}

		if sep {
			if len(cur) > 0 {
				solveP2(cur)
				cur = cur[:0]
			}
			continue
		}

		cur = append(cur, c)
	}
	if len(cur) > 0 {
		solveP2(cur)
	}

	return Input{p1, p2}
}

func Part1(input Input) uint64 {
	return input.p1
}

func Part2(input Input) uint64 {
	return input.p2
}
