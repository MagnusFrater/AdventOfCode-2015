package circuit

import "fmt"

type formula interface {
	Formulate()
}

// set is a vertex in the graph.
// e.g. 123 -> x
// e.g. a -> b
type set struct {
	Signal     uint16 // only set if input is a signal
	InputWire  *wire  // only set if input is a wire
	TargetWire *wire
}

// not is a vertex in the graph.
// e.g. NOT a -> b
type not struct {
	InputWire  *wire
	TargetWire *wire
}

// and is a vertex in the graph.
// e.g. a AND b -> c
// e.g. 1 AND b -> c
type and struct {
	Signal1                uint16
	InputWire1, InputWire2 *wire
	TargetWire             *wire
}

// or is a vertex in the graph
// e.g. a OR b -> c
type or struct {
	InputWire1, InputWire2 *wire
	TargetWire             *wire
}

// lshift is a vertex in the graph.
// e.g. a LSHIFT 2 -> b
type lshift struct {
	InputWire   *wire
	ShiftAmount uint16
	TargetWire  *wire
}

// rshift is a vertex in the graph.
// e.g. a RSHIFT 2 -> b
type rshift struct {
	InputWire   *wire
	ShiftAmount uint16
	TargetWire  *wire
}

func (s set) Formulate() {
	if s.InputWire != nil {
		// input is wire
		if !s.InputWire.Simulated {
			// input wire is not ready (doesn't have a value to work on yet)
			fmt.Printf("SET: %v -> %v (skipping; input wire not simulated)\n", s.InputWire.Name, s.TargetWire.Name)
			return
		}

		// input wire does have a value to work on
		s.TargetWire.Signal = s.InputWire.Signal
		fmt.Printf("SET: %v (%v) -> %v (%v)\n",
			s.InputWire.Name, s.InputWire.Signal, s.TargetWire.Name, s.TargetWire.Signal)
	} else {
		// input is signal
		s.TargetWire.Signal = s.Signal
		fmt.Printf("SET: %v -> %v (%v)\n", s.Signal, s.TargetWire.Name, s.TargetWire.Signal)
	}

	// enter the target wire to continue doing work down further into the circuit
	s.TargetWire.simulate()
}

func (n not) Formulate() {
	if !n.InputWire.Simulated {
		// input wire is not ready (doesn't have a value to work on yet)
		fmt.Printf("NOT: %v -> %v (skipping; input wire not simulated)\n", n.InputWire.Name, n.TargetWire.Name)
		return
	}

	n.TargetWire.Signal = ^n.InputWire.Signal
	fmt.Printf("NOT: %v (%v) -> %v (%v)\n",
		n.InputWire.Name, n.InputWire.Signal, n.TargetWire.Name, n.TargetWire.Signal)

	n.TargetWire.simulate()
}

func (a and) Formulate() {
	fmt.Println(a)
	if !a.InputWire2.Simulated {
		// input wire is not ready (doesn't have a value to work on yet)
		fmt.Printf("AND: %v & %v -> %v (skipping; input wire %v not simulated)\n",
			a.InputWire1.Name, a.InputWire2.Name, a.TargetWire.Name, a.InputWire2.Name)
		return
	}

	if a.InputWire1 != nil {
		// leading input is wire
		if !a.InputWire1.Simulated {
			// input wire is not ready (doesn't have a value to work on yet)
			fmt.Printf("AND: %v & %v -> %v (skipping; input wire 1 not simulated)\n",
				a.InputWire1.Name, a.InputWire2.Name, a.TargetWire.Name)
			return
		}

		a.TargetWire.Signal = a.InputWire1.Signal & a.InputWire2.Signal
		fmt.Printf("AND: %v (%v) & %v (%v) -> %v (%v)\n",
			a.InputWire1.Name, a.InputWire1.Signal, a.InputWire2.Name, a.InputWire2.Signal,
			a.TargetWire.Name, a.TargetWire.Signal)
	} else {
		// leading input is signal
		a.TargetWire.Signal = a.Signal1 & a.InputWire2.Signal
		fmt.Printf("AND: %v & %v (%v) -> %v (%v)\n",
			a.Signal1, a.InputWire2.Name, a.InputWire2.Signal, a.TargetWire.Name, a.TargetWire.Signal)
	}

	a.TargetWire.simulate()
}

func (o or) Formulate() {
	if !o.InputWire1.Simulated || !o.InputWire2.Simulated {
		// input wire(s) is not ready (doesn't have a value to work on yet)
		var unreadyWire = o.InputWire1.Name
		if !o.InputWire2.Simulated {
			unreadyWire = o.InputWire2.Name
		}
		fmt.Printf("OR: %v | %v -> %v (skipping; input wire %v not simulated)\n",
			o.InputWire1.Name, o.InputWire2.Name, o.TargetWire.Name, unreadyWire)
		return
	}

	o.TargetWire.Signal = o.InputWire1.Signal | o.InputWire2.Signal
	fmt.Printf("OR: %v (%v) | %v (%v) -> %v (%v)\n",
		o.InputWire1.Name, o.InputWire1.Signal, o.InputWire2.Name, o.InputWire2.Signal,
		o.TargetWire.Name, o.TargetWire.Signal)

	o.TargetWire.simulate()
}

func (ls lshift) Formulate() {
	if !ls.InputWire.Simulated {
		// input wire is not ready (doesn't have a value to work on yet)
		fmt.Printf("LSHIFT: %v & %v -> %v (skipping; input wire not simulated)\n",
			ls.InputWire.Name, ls.ShiftAmount, ls.TargetWire.Name)
		return
	}

	ls.TargetWire.Signal = ls.InputWire.Signal << ls.ShiftAmount
	fmt.Printf("LSHIFT: %v (%v) & %v -> %v (%v)\n",
		ls.InputWire.Name, ls.InputWire.Signal, ls.ShiftAmount, ls.TargetWire.Name, ls.TargetWire.Signal)

	ls.TargetWire.simulate()
}

func (rs rshift) Formulate() {
	if !rs.InputWire.Simulated {
		// input wire is not ready (doesn't have a value to work on yet)
		fmt.Printf("RSHIFT: %v & %v -> %v (skipping; input wire not simulated)\n",
			rs.InputWire.Name, rs.ShiftAmount, rs.TargetWire.Name)
		return
	}

	rs.TargetWire.Signal = rs.InputWire.Signal >> rs.ShiftAmount
	fmt.Printf("RSHIFT: %v (%v) & %v -> %v (%v)\n",
		rs.InputWire.Name, rs.InputWire.Signal, rs.ShiftAmount, rs.TargetWire.Name, rs.TargetWire.Signal)

	rs.TargetWire.simulate()
}
