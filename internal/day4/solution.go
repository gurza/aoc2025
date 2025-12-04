package day4

type Input struct {
	p1 uint64
	p2 uint64
}

func Parse(input []string) Input {
	var cnt1, cnt2 uint64

	rows := len(input)
	cols := len(input[0])

	grid := make([][]byte, rows)
	for i := range grid {
		grid[i] = make([]byte, cols)
		copy(grid[i], input[i])
	}

	isAt := func(r, c int) bool {
		return r >= 0 && r < rows && c >= 0 && c < cols && grid[r][c] == '@'
	}

	dirs := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	firstPass := true
	adj := 0
	for {
		var toRemove [][2]int

		for r := range rows {
			for c := range cols {
				if grid[r][c] != '@' {
					continue
				}

				adj = 0
				for _, dir := range dirs {
					cr, cc := r+dir[0], c+dir[1]
					if isAt(cr, cc) {
						adj += 1
					}
				}
				if adj < 4 {
					toRemove = append(toRemove, [2]int{r, c})
				}
			}
		}

		if len(toRemove) == 0 {
			break
		}

		if firstPass {
			cnt1 = uint64(len(toRemove))
			firstPass = false
		}

		for _, pos := range toRemove {
			grid[pos[0]][pos[1]] = '.'
		}
		cnt2 += uint64(len(toRemove))
	}

	return Input{cnt1, cnt2}
}

func Part1(input Input) uint64 {
	return input.p1
}

func Part2(input Input) uint64 {
	return input.p2
}
