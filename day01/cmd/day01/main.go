package main

import (
	"fmt"

	"github.com/MagnusFrater/AdventOfCode-2015/day01"
	"github.com/MagnusFrater/AdventOfCode-2015/internal/load"
)

func main() {
	var input = load.Input("1")
	fmt.Println("Part 1: floor", day01.Part1(input))
	fmt.Println("Part 2: index", day01.Part2(input))
}
