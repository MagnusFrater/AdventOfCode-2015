package main

import (
	"fmt"
	"strings"

	"github.com/MagnusFrater/AdventOfCode-2015/day05"
	"github.com/MagnusFrater/AdventOfCode-2015/internal/load"
)

func main() {
	input := strings.Split(load.Input("05"), "\n")
	fmt.Println("Part 1:", day05.Part1(input), "nice words")
	fmt.Println("Part 2:", day05.Part2(input), "nice words")
}
