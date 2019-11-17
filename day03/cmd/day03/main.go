package main

import (
	"fmt"

	"github.com/MagnusFrater/AdventOfCode-2015/day03"
	"github.com/MagnusFrater/AdventOfCode-2015/internal/load"
)

func main() {
	fmt.Printf("Part 1: %v houses\n", day03.Part1(load.Input("3")))
	fmt.Printf("Part 2: %v houses\n", day03.Part2(load.Input("3")))
}
