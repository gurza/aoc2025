package day4

type Input struct {
	p1 uint64
	p2 uint64
}

func Parse(input []string) Input {
	var cnt1, cnt2 uint64

	rows := len(input)
	cols := len(input[0])

	isAt := func(r, c int) bool {
		return r >= 0 && r < rows && c >= 0 && c < cols && input[r][c] == '@'
	}

	dirs := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	adjCnt := 0
	for r := range rows {
		for c := range cols {
			if input[r][c] != '@' {
				continue
			}

			adjCnt = 0
			for _, dir := range dirs {
				cr, cc := r+dir[0], c+dir[1]
				if isAt(cr, cc) {
					adjCnt += 1
				}
			}
			if adjCnt < 4 {
				cnt1 += 1
			}
		}
	}

	return Input{cnt1, cnt2}
}

func Part1(input Input) uint64 {
	return input.p1
}

func Part2(input Input) uint64 {
	return input.p2
}
