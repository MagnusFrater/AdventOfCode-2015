package main

import (
	"fmt"
	"strings"

	"github.com/MagnusFrater/AdventOfCode-2015/day08"
	"github.com/MagnusFrater/AdventOfCode-2015/internal/load"
)

func main() {
	input := strings.Split(load.Input("08"), "\n")
	fmt.Printf("Part 1: %v characters\n", day08.Part1(input))
	fmt.Printf("Part 2: %v characters\n", day08.Part2(input))
}
