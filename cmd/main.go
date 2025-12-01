package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gurza/aoc2025/internal/day1"
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
		data := day1.Parse(lines)
		fmt.Printf("Part 1: %d\n", day1.Part1(data))
		fmt.Printf("Part 2: %d\n", day1.Part2(data))
	default:
		fmt.Fprintf(os.Stderr, "Day %d not implemented yet.\n", *day)
	}
}
