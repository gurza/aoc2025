package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gurza/aoc2025/internal/day1"
	"github.com/gurza/aoc2025/internal/day11"
	"github.com/gurza/aoc2025/internal/day2"
	"github.com/gurza/aoc2025/internal/day3"
	"github.com/gurza/aoc2025/internal/day4"
	"github.com/gurza/aoc2025/internal/day5"
	"github.com/gurza/aoc2025/internal/day6"
	"github.com/gurza/aoc2025/internal/day7"
	"github.com/gurza/aoc2025/internal/day8"
	"github.com/gurza/aoc2025/internal/day9"
)

func main() {
	day := flag.Int("day", 1, "Which Advent of Code 2025 day to run")
	flag.Parse()

	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
	lines := strings.Split(string(input), "\n")
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	switch *day {
	case 1:
		input := day1.Parse(lines)
		fmt.Printf("Part 1: %d\n", day1.Part1(input))
		fmt.Printf("Part 2: %d\n", day1.Part2(input))
	case 2:
		input := day2.Parse(lines)
		fmt.Printf("Part 1: %d\n", day2.Part1(input))
		fmt.Printf("Part 2: %d\n", day2.Part2(input))
	case 3:
		input := day3.Parse(lines)
		fmt.Printf("Part 1: %d\n", day3.Part1(input))
		fmt.Printf("Part 2: %d\n", day3.Part2(input))
	case 4:
		input := day4.Parse(lines)
		fmt.Printf("Part 1: %d\n", day4.Part1(input))
		fmt.Printf("Part 2: %d\n", day4.Part2(input))
	case 5:
		input := day5.Parse(lines)
		fmt.Printf("Part 1: %d\n", day5.Part1(input))
		fmt.Printf("Part 2: %d\n", day5.Part2(input))
	case 6:
		input := day6.Parse(lines)
		fmt.Printf("Part 1: %d\n", day6.Part1(input))
		fmt.Printf("Part 2: %d\n", day6.Part2(input))
	case 7:
		input := day7.Parse(lines)
		fmt.Printf("Part 1: %d\n", day7.Part1(input))
		fmt.Printf("Part 2: %d\n", day7.Part2(input))
	case 8:
		input := day8.Parse(lines)
		fmt.Printf("Part 1: %d\n", day8.Part1(input))
		fmt.Printf("Part 2: %d\n", day8.Part2(input))
	case 9:
		input := day9.Parse(lines)
		fmt.Printf("Part 1: %d\n", day9.Part1(input))
		fmt.Printf("Part 2: %d\n", day9.Part2(input))
	case 11:
		input := day11.Parse(lines)
		fmt.Printf("Part 1: %d\n", day11.Part1(input))
		fmt.Printf("Part 2: %d\n", day11.Part2(input))
	default:
		fmt.Fprintf(os.Stderr, "Day %d not implemented yet.\n", *day)
	}
}
