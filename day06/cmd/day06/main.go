package main

import (
	"fmt"
	"strings"

	"github.com/MagnusFrater/AdventOfCode-2015/day06"
	"github.com/MagnusFrater/AdventOfCode-2015/internal/load"
)

func main() {
	input := strings.Split(load.Input("06"), "\n")
	fmt.Printf("Part 1: %v lights on\n", day06.Part1(input))
	fmt.Printf("Part 2: %v total brightness\n", day06.Part2(input))
}
