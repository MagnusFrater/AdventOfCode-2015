package main

import (
	"fmt"
	"strings"

	"github.com/MagnusFrater/AdventOfCode-2015/day02"
	"github.com/MagnusFrater/AdventOfCode-2015/internal/load"
)

func main() {
	var input = strings.Split(load.Input("2"), "\n")
	fmt.Printf("Part 1: %v ft^2\n", day02.Part1(input))
	fmt.Printf("Part 2: %v ft\n", day02.Part2(input))
}
