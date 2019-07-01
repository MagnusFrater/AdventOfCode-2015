package day05

import (
	"testing"

	"github.com/MagnusFrater/AdventOfCode-2015/internal/test"
)

func TestPart1(t *testing.T) {
	a := test.NewAssert(t)

	var testCases = []test.Case{
		{Input: "ugknbfddgicrmopn", Expected: 1},
		{Input: "aaa", Expected: 1},
		{Input: "jchzalrnumimnmhp", Expected: 0},
		{Input: "haegwjzuvuyypxyu", Expected: 0},
		{Input: "dvszwmarrgswjxmb", Expected: 0},
	}

	for _, tc := range testCases {
		input := []string{tc.Input.(string)}
		a.Equals(tc, Part1(input))
	}
}

func TestPart2(t *testing.T) {
	a := test.NewAssert(t)

	var testCases = []test.Case{
		{Input: "qjhvhtzxzqqjkmpb", Expected: 1},
		{Input: "xxyxx", Expected: 1},
		{Input: "uurcxstgmygtbstg", Expected: 0},
		{Input: "ieodomkazucvgmuy", Expected: 0},
	}

	for _, tc := range testCases {
		input := []string{tc.Input.(string)}
		a.Equals(tc, Part2(input))
	}
}
