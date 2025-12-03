package day3

type Input struct {
	p1 uint64
	p2 uint64
}

func Parse(input []string) Input {
	var sum1, sum2 uint64

	for _, s := range input {
		left := s[0] - '0'
		max := uint64(0)

		for i := 1; i < len(s); i++ {
			cur := s[i] - '0'
			cand := uint64(10*left + cur)
			if cand > max {
				max = cand
			}
			if cur > left {
				left = cur
			}
		}

		sum1 += max
	}

	return Input{sum1, sum2}
}

func Part1(input Input) uint64 {
	return input.p1
}

func Part2(input Input) uint64 {
	return input.p2
}
