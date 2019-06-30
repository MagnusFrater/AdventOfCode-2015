package day01

import (
	"testing"

	"github.com/MagnusFrater/AdventOfCode-2015/internal/test"
)

func TestPart1(t *testing.T) {
	a := test.NewAssert(t)

	var testCases = []test.Case{
		{Input: "(())", Expected: 0},
		{Input: "()()", Expected: 0},
		{Input: "(((", Expected: 3},
		{Input: "(()(()(", Expected: 3},
		{Input: "))(((((", Expected: 3},
		{Input: "())", Expected: -1},
		{Input: ")))", Expected: -3},
		{Input: ")())())", Expected: -3},
	}

	for _, tc := range testCases {
		a.Equals(tc, Part1(tc.Input.(string)))
	}
}

func TestPart2(t *testing.T) {
	a := test.NewAssert(t)

	testCases := []test.Case{
		{Input: ")", Expected: 1},
		{Input: "()())", Expected: 5},
	}

	for _, tc := range testCases {
		a.Equals(tc, Part2(tc.Input.(string)))
	}
}
