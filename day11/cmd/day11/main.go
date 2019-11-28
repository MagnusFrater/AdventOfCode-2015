package main

import (
	"fmt"

	"github.com/MagnusFrater/AdventOfCode-2015/day11"
	"github.com/MagnusFrater/AdventOfCode-2015/internal/load"
)

func main() {
	input := load.Input("11")
	fmt.Printf("Part 1: %v\n", day11.Part1(input))
	fmt.Printf("Part 2: %v\n", day11.Part2(input))
}
