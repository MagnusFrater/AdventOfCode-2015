package main

import (
	"fmt"
	"strings"

	"github.com/MagnusFrater/AdventOfCode-2015/day06"
	"github.com/MagnusFrater/AdventOfCode-2015/internal/load"
)

func main() {
	input := strings.Split(load.Input("06"), "\n")
	fmt.Println("Part 1:", day06.Part1(input), "lights on")
	fmt.Println("Part 2:", day06.Part2(input), "total brightness")
}
