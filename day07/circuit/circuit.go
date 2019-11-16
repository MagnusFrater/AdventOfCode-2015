package circuit

import (
	"fmt"
	"strings"
)

// Circuit is a directed acyclic graph.
type Circuit struct {
	wires map[string]*wire
	roots []*set
}

// New creates a new circuit.
func New() Circuit {
	return Circuit{
		wires: make(map[string]*wire),
		roots: []*set{},
	}
}

// AddInstructions loads the given instruction list into the circuit.
func (c *Circuit) AddInstructions(instructions []string) {
	for _, instruction := range instructions {
		c.AddInstruction(instruction)
	}
}

// AddInstruction loads an individual instruction into the circuit.
func (c *Circuit) AddInstruction(instruction string) {
	split := strings.Fields(instruction)

	if strings.Contains(instruction, "NOT") {
		// create formula vertex
		var not = not{
			InputWire:  c.createWireIfNotExist(split[1]),
			TargetWire: c.createWireIfNotExist(split[3]),
		}

		// connect input->formula
		not.InputWire.addConnection(not)
	} else if strings.Contains(instruction, "AND") {
		// create formula vertex
		var and = and{
			InputWire2: c.createWireIfNotExist(split[2]),
			TargetWire: c.createWireIfNotExist(split[4]),
		}

		var input = split[0]
		signal, err := getUint16(input)
		if err != nil {
			// is wire
			// connect wire->formula
			and.InputWire1 = c.createWireIfNotExist(input)
			and.InputWire1.addConnection(and)
		} else {
			// is signal
			and.Signal1 = signal
		}

		// connect wire->formula
		and.InputWire2.addConnection(and)
	} else if strings.Contains(instruction, "OR") {
		// create formula vertex
		var or = or{
			InputWire1: c.createWireIfNotExist(split[0]),
			InputWire2: c.createWireIfNotExist(split[2]),
			TargetWire: c.createWireIfNotExist(split[4]),
		}

		// connect inputs->formula
		or.InputWire1.addConnection(or)
		or.InputWire2.addConnection(or)
	} else if strings.Contains(instruction, "LSHIFT") {
		shiftAmount, err := getUint16(split[2])
		if err != nil {
			panic(err)
		}

		// create formula vertex
		var lshift = lshift{
			InputWire:   c.createWireIfNotExist(split[0]),
			ShiftAmount: shiftAmount,
			TargetWire:  c.createWireIfNotExist(split[4]),
		}

		// connect input->formula
		lshift.InputWire.addConnection(lshift)
	} else if strings.Contains(instruction, "RSHIFT") {
		shiftAmount, err := getUint16(split[2])
		if err != nil {
			panic(err)
		}

		// create formula vertex
		var rshift = rshift{
			InputWire:   c.createWireIfNotExist(split[0]),
			ShiftAmount: shiftAmount,
			TargetWire:  c.createWireIfNotExist(split[4]),
		}

		// connect input->formula
		rshift.InputWire.addConnection(rshift)
	} else { // SET
		// create formula vertex
		var set = set{
			TargetWire: c.createWireIfNotExist(split[2]),
		}

		// check if input is signal or wire
		var input = split[0]
		signal, err := getUint16(input)
		if err != nil {
			// is wire
			var inputWire = c.createWireIfNotExist(input)

			// connect wire->formula
			set.InputWire = inputWire
			inputWire.addConnection(set)
		} else {
			// is signal
			set.Signal = signal
			c.addRoot(set)
		}
	}
}

func (c *Circuit) addRoot(s set) {
	c.roots = append(c.roots, &s)
}

func (c *Circuit) createWireIfNotExist(name string) *wire {
	if _, prs := c.wires[name]; !prs {
		c.wires[name] = &wire{
			Name: name,
		}
	}
	return c.wires[name]
}

// Simulate propagates the signals through the circuit.
func (c Circuit) Simulate() {
	for _, root := range c.roots {
		root.Formulate()
	}
}

// GetSignal returns the signal of the given wire.
func (c Circuit) GetSignal(wireName string) uint16 {
	return c.wires[wireName].Signal
}

// Print displays all wires and their values.
func (c Circuit) Print() {
	fmt.Println("Wires:")
	for name, wire := range c.wires {
		fmt.Printf("%v: %v\n", name, wire.Signal)
	}
}
