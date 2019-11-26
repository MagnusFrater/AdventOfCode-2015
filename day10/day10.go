package day10

import (
	"strconv"
)

// Part1 returns the output for Day 10 part 1.
func Part1(input string) int {
	for i := 0; i < 40; i++ {
		input = lookAndSay(input)
	}
	return len(input)
}

// Part2 returns the output for Day 10 part 2.
func Part2(input string) int {
	for i := 0; i < 50; i++ {
		input = lookAndSay(input)
	}
	return len(input)
}

func lookAndSay(n string) string {
	var builder string
	var current rune
	var count int

	for _, c := range n {
		if c != current {
			if count > 0 {
				builder += strconv.Itoa(count) + string(current)
			}

			current = c
			count = 1
		} else {
			count++
		}
	}

	builder += strconv.Itoa(count) + string(current)

	return builder
}
