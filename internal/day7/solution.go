package day7

import "strings"

type Input struct {
	p1 uint64
	p2 uint64
}

func Parse(input []string) Input {
	var p1, p2 uint64

	rows := len(input)
	cols := len(input[0])

	sRow, sCol := -1, -1
	for r, row := range input {
		if c := strings.IndexByte(row, 'S'); c != -1 {
			sRow, sCol = r, c
			break
		}
	}

	seen := make([][]bool, rows)
	for i := range seen {
		seen[i] = make([]bool, cols)
	}

	beams := make([]uint64, cols)
	beams[sCol] = 1 // one beam starting at S going down
	for r := sRow + 1; r < rows; r++ {
		row := input[r]
		newBeams := make([]uint64, cols)

		for c, b := range beams {
			if b == 0 {
				continue
			}
			switch row[c] {
			case '.':
				newBeams[c] += b
			case '^':
				if !seen[r][c] {
					seen[r][c] = true
					p1++
				}

				if c > 0 {
					newBeams[c-1] += b
				}
				if c+1 < cols {
					newBeams[c+1] += b
				}
			}
		}
		beams = newBeams
	}

	return Input{p1, p2}
}

func Part1(input Input) uint64 {
	return input.p1
}

func Part2(input Input) uint64 {
	return input.p2
}
