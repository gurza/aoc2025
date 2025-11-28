package main

import (
	"flag"
	"fmt"

	day01 "github.com/gurza/aoc2025/internal/day1"
)

func main() {
	day := flag.Int("day", 1, "Which Advent of Code 2025 day to run")
	flag.Parse()

	switch *day {
	case 1:
		day01.Solve()
	default:
		fmt.Printf("Day %d not implemented yet.\n", *day)
	}
}
