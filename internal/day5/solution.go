package day5

import (
	"math"
	"sort"
)

type Input struct {
	p1 uint64
	p2 uint64
}

// rangePair represents an inclusive [start, end] interval.
type rangePair struct {
	start int
	end   int
}

func Parse(input []string) Input {
	var p1, p2 uint64

	afterBlank := false
	for _, line := range input {
		if line == "" {
			afterBlank = true
			continue
		}
		if !afterBlank {
			// handle range

			continue
		}

		// handle ID
	}

	return Input{p1, p2}
}

// mergeRanges sorts ranges in-place by start and merges overlapping or
// directly adjacent inclusive intervals.
//
// Warning! The input slice is reordered.
func mergeRanges(ranges []rangePair) []rangePair {
	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i].start == ranges[j].start {
			return ranges[i].end < ranges[j].end
		}
		return ranges[i].start < ranges[j].start
	})

	n := len(ranges)
	out := ranges[:0]
	cur := ranges[0]
	for i := 1; i < n; i++ {
		r := ranges[i]
		ov := r.start <= cur.end
		adj := !ov && cur.end != math.MaxInt && r.start == cur.end+1
		if ov || adj {
			if r.end > cur.end {
				cur.end = r.end
			}
			continue
		}
		out = append(out, cur)
		cur = r
	}
	out = append(out, cur)

	return out
}

func Part1(input Input) uint64 {
	return input.p1
}

func Part2(input Input) uint64 {
	return input.p2
}
