package main

import (
	"fmt"

	"github.com/MagnusFrater/AdventOfCode-2015/day01"
	"github.com/MagnusFrater/AdventOfCode-2015/load"
)

func main() {
	var input = load.InputFileAsString("01")
	fmt.Println(day01.Part1(input))
	fmt.Println(day01.Part2(input))
}
