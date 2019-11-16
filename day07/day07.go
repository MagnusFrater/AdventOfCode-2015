package day07

import (
	"github.com/MagnusFrater/AdventOfCode-2015/day07/circuit"
)

// Part1 returns the output for Day 07 part 1.
func Part1(instructions []string) uint16 {
	var circuit = circuit.New()

	circuit.AddInstructions(instructions)
	circuit.Simulate()

	return circuit.GetSignal("a")
}

// Part2 returns the output for Day 07 part2.
func Part2(instructions []string) uint16 {
	var circuit = circuit.New()

	circuit.AddInstructions(instructions)
	circuit.Simulate()

	circuit.SetRoot(circuit.GetSignal("a"), "b")
	circuit.Simulate()

	return circuit.GetSignal("a")
}
