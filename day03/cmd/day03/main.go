package main

import (
	"fmt"

	"github.com/MagnusFrater/AdventOfCode-2015/day03"
	"github.com/MagnusFrater/AdventOfCode-2015/internal/load"
)

func main() {
	fmt.Println("Part 1:", day03.Part1(load.Input("3")), "houses")
	fmt.Println("Part 2:", day03.Part2(load.Input("3")), "houses")
}
