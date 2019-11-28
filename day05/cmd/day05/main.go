package main

import (
	"fmt"
	"strings"

	"github.com/MagnusFrater/AdventOfCode-2015/day05"
	"github.com/MagnusFrater/AdventOfCode-2015/internal/load"
)

func main() {
	input := strings.Split(load.Input("05"), "\n")
	fmt.Printf("Part 1: %v nice words\n", day05.Part1(input))
	fmt.Printf("Part 2: %v nice words\n", day05.Part2(input))
}
