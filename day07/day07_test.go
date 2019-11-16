package day07

import (
	"testing"
)

func TestAddInstructions(t *testing.T) {
	type Case struct {
		Instructions    []string
		ExpectedSignals map[string]uint16
	}

	// var tc = Case{
	// 	[]string{
	// 		"123 -> x",
	// 		"456 -> y",
	// 		"x AND y -> d",
	// 		"x OR y -> e",
	// 		"x LSHIFT 2 -> f",
	// 		"y RSHIFT 2 -> g",
	// 		"NOT x -> h",
	// 		"NOT y -> i",
	// 	},
	// 	map[string]uint16{
	// 		"d": 72,
	// 		"e": 507,
	// 		"f": 492,
	// 		"g": 114,
	// 		"h": 65412,
	// 		"i": 65079,
	// 		"x": 123,
	// 		"y": 456,
	// 	},
	// }

	// var circuit circuit
	// circuit.addInstructions(tc.Instructions)

	// for wireName, signal := range circuit.Wires {
	// 	fmt.Printf("%v: %v\n", wireName, signal)
	// }

	// for wireName, signal := range tc.ExpectedSignals {
	// 	wire, prs := circuit.Wires[wireName]

	// 	if !prs {
	// 		t.Errorf("Wire not found: %v\n", wireName)
	// 		continue
	// 	}

	// 	if signal != wire.Signal {
	// 		t.Errorf("Wire %v has incorrect signal:\tExpected: %v\tActual:%v\n", wireName, signal, wire.Signal)
	// 	}
	// }
}
