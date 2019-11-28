package main

import (
	"fmt"

	"github.com/MagnusFrater/AdventOfCode-2015/day01"
	"github.com/MagnusFrater/AdventOfCode-2015/internal/load"
)

func main() {
	var input = load.Input("1")
	fmt.Printf("Part 1: floor %v\n", day01.Part1(input))
	fmt.Printf("Part 2: index %v\n", day01.Part2(input))
}
