package day09

import (
	"strings"
	"testing"

	"github.com/MagnusFrater/AdventOfCode-2015/internal/load"
)

type testCase struct {
	Edges            []string
	ExpectedDistance int
}

func TestPart1(t *testing.T) {
	var tc = testCase{
		Edges: []string{
			"London to Dublin = 464",
			"London to Belfast = 518",
			"Dublin to Belfast = 141",
		},
		ExpectedDistance: 605,
	}

	var distance = Part1(tc.Edges)
	if distance != tc.ExpectedDistance {
		t.Errorf("Expected: %v\tActual: %v\n", tc.ExpectedDistance, distance)
	}
}

func TestPart1_solution(t *testing.T) {
	var tc = testCase{
		Edges:            strings.Split(load.InputFileAsString("input09.txt"), "\n"),
		ExpectedDistance: 251,
	}

	var distance = Part1(tc.Edges)
	if distance != tc.ExpectedDistance {
		t.Errorf("Expected: %v\tActual: %v\n", tc.ExpectedDistance, distance)
	}
}

func TestPart2(t *testing.T) {
	var tc = testCase{
		Edges: []string{
			"London to Dublin = 464",
			"London to Belfast = 518",
			"Dublin to Belfast = 141",
		},
		ExpectedDistance: 982,
	}

	var distance = Part2(tc.Edges)
	if distance != tc.ExpectedDistance {
		t.Errorf("Expected: %v\tActual: %v\n", tc.ExpectedDistance, distance)
	}
}

func TestPart2_solution(t *testing.T) {
	var tc = testCase{
		Edges:            strings.Split(load.InputFileAsString("input09.txt"), "\n"),
		ExpectedDistance: 898,
	}

	var distance = Part2(tc.Edges)
	if distance != tc.ExpectedDistance {
		t.Errorf("Expected: %v\tActual: %v\n", tc.ExpectedDistance, distance)
	}
}
