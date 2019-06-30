package main

import (
	"fmt"
	"strings"

	"github.com/MagnusFrater/AdventOfCode-2015/day02"
	"github.com/MagnusFrater/AdventOfCode-2015/internal/load"
)

func main() {
	var input = strings.Split(load.Input("2"), "\n")
	fmt.Println("Part 1:", day02.Part1(input), "ft^2")
	fmt.Println("Part 2:", day02.Part2(input), "ft")
}
