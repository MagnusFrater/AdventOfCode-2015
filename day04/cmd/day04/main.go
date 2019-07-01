package main

import (
	"fmt"

	"github.com/MagnusFrater/AdventOfCode-2015/day04"
	"github.com/MagnusFrater/AdventOfCode-2015/internal/load"
)

func main() {
	input := load.Input("04")
	fmt.Println("Part 1:", day04.Part1(input, 1))
	fmt.Println("Part 2:", day04.Part2(input, 1))
}
