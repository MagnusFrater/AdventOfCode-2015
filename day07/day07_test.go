package day07

import (
	"strings"
	"testing"

	"github.com/MagnusFrater/AdventOfCode-2015/internal/load"
)

type testCase struct {
	Instructions   []string
	ExpectedSignal uint16
}

func TestPart1_solution(t *testing.T) {
	var tc = testCase{
		Instructions:   strings.Split(load.InputFileAsString("input07.txt"), "\n"),
		ExpectedSignal: 956,
	}

	var signal = Part1(tc.Instructions)
	if signal != tc.ExpectedSignal {
		t.Errorf("Expected: %v\tActual: %v\n", tc.ExpectedSignal, signal)
	}
}

func TestPart2_solution(t *testing.T) {
	var tc = testCase{
		Instructions:   strings.Split(load.InputFileAsString("input07.txt"), "\n"),
		ExpectedSignal: 40149,
	}

	var signal = Part2(tc.Instructions)
	if signal != tc.ExpectedSignal {
		t.Errorf("Expected: %v\tActual: %v\n", tc.ExpectedSignal, signal)
	}
}
