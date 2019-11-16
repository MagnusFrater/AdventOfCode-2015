package circuit

import (
	"testing"
)

func TestAddInstructions(t *testing.T) {
	type TestCase struct {
		Instructions    []string
		ExpectedSignals map[string]uint16
	}

	var tc = TestCase{
		Instructions: []string{
			"123 -> x",
			"456 -> y",
			"x AND y -> d",
			"x OR y -> e",
			"x LSHIFT 2 -> f",
			"y RSHIFT 2 -> g",
			"NOT x -> h",
			"NOT y -> i",
		},
		ExpectedSignals: map[string]uint16{
			"d": 72,
			"e": 507,
			"f": 492,
			"g": 114,
			"h": 65412,
			"i": 65079,
			"x": 123,
			"y": 456,
		},
	}

	var circuit = New()
	circuit.AddInstructions(tc.Instructions)
	circuit.Simulate()

	for wireName, signal := range tc.ExpectedSignals {
		wire, prs := circuit.wires[wireName]

		if !prs {
			t.Errorf("Wire not found: %v\n", wireName)
			continue
		}

		if signal != wire.Signal {
			t.Errorf("Wire %v has incorrect signal:\tExpected: %v\tActual:%v\n", wireName, signal, wire.Signal)
		}
	}
}
