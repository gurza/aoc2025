package day3

type Input struct {
	p1 uint64
	p2 uint64
}

func Parse(input []string) Input {
	const n = 12
	var sum1, sum2 uint64

	for _, s := range input {
		sLen := len(s)

		// for two batteries
		var max2 uint64
		left2 := s[0] - '0'

		// for n batteries
		var stack [n]uint8
		stackPos := 0
		toDrop := sLen - n

		// single pass over s
		for i := range sLen {
			cur := s[i] - '0'

			// n batteries
			for stackPos > 0 && stack[stackPos-1] < cur && toDrop > 0 {
				stackPos--
				toDrop--
			}
			if stackPos < n {
				stack[stackPos] = cur
				stackPos++
			} else {
				toDrop--
			}

			// two batteries
			cand := uint64(10*left2 + cur)
			if cand > max2 {
				max2 = cand
			}
			if cur > left2 {
				left2 = cur
			}
		}

		sum1 += max2
		sum2 += stackToNumber(stack[:])
	}

	return Input{sum1, sum2}
}

func stackToNumber(stack []uint8) uint64 {
	var res uint64
	n := len(stack)
	for i := range n {
		res = res*10 + uint64(stack[i])
	}
	return res
}

func Part1(input Input) uint64 {
	return input.p1
}

func Part2(input Input) uint64 {
	return input.p2
}
