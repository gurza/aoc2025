package day5

import (
	"math"
	"sort"
	"strconv"
	"strings"
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
	var rawRanges []rangePair
	var mergedRanges []rangePair
	for _, line := range input {
		if line == "" {
			mergedRanges = mergeRanges(rawRanges)
			afterBlank = true
			continue
		}
		if !afterBlank {
			// handle range, not safety
			pp := strings.SplitN(line, "-", 2)
			start, _ := strconv.Atoi(pp[0])
			end, _ := strconv.Atoi(pp[1])
			rawRanges = append(rawRanges, rangePair{start, end})
			continue
		}

		// handle ID, not safety too
		id, _ := strconv.Atoi(line)
		if isFresh(id, mergedRanges) {
			p1++
		}
	}
	for _, r := range mergedRanges {
		p2 += uint64(r.end - r.start + 1)
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

func isFresh(id int, ranges []rangePair) bool {
	lo, hi := 0, len(ranges)
	for lo < hi {
		mid := (lo + hi) / 2
		r := ranges[mid]
		if id < r.start {
			hi = mid
		} else if id > r.end {
			lo = mid + 1
		} else {
			return true
		}
	}
	return false
}

func Part1(input Input) uint64 {
	return input.p1
}

func Part2(input Input) uint64 {
	return input.p2
}
