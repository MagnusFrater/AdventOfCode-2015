package main

import (
	"fmt"
	"strings"

	"github.com/MagnusFrater/AdventOfCode-2015/day07"
	"github.com/MagnusFrater/AdventOfCode-2015/internal/load"
)

func main() {
	input := strings.Split(load.Input("07"), "\n")
	fmt.Println("Part 1:", day07.Part1(input))
	fmt.Println("Part 2:", day07.Part2(input))
}
