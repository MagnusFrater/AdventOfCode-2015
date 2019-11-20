package main

import (
	"fmt"
	"strings"

	"github.com/MagnusFrater/AdventOfCode-2015/day09"
	"github.com/MagnusFrater/AdventOfCode-2015/internal/load"
)

func main() {
	input := strings.Split(load.Input("09"), "\n")
	fmt.Printf("Part 1: %v\n", day09.Part1(input))
	fmt.Printf("Part 2: %v\n", day09.Part2(input))
}
