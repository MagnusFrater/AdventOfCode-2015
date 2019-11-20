package circuit

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
			return
		}

		// input wire does have a value to work on
		s.TargetWire.Signal = s.InputWire.Signal
	} else {
		// input is signal
		s.TargetWire.Signal = s.Signal
	}

	// enter the target wire to continue doing work down further into the circuit
	s.TargetWire.simulate()
}

func (n not) Formulate() {
	if !n.InputWire.Simulated {
		// input wire is not ready (doesn't have a value to work on yet)
		return
	}

	n.TargetWire.Signal = ^n.InputWire.Signal
	n.TargetWire.simulate()
}

func (a and) Formulate() {
	if !a.InputWire2.Simulated {
		// input wire is not ready (doesn't have a value to work on yet)
		return
	}

	if a.InputWire1 != nil {
		// leading input is wire
		if !a.InputWire1.Simulated {
			// input wire is not ready (doesn't have a value to work on yet)
			return
		}

		a.TargetWire.Signal = a.InputWire1.Signal & a.InputWire2.Signal
	} else {
		// leading input is signal
		a.TargetWire.Signal = a.Signal1 & a.InputWire2.Signal
	}

	a.TargetWire.simulate()
}

func (o or) Formulate() {
	if !o.InputWire1.Simulated || !o.InputWire2.Simulated {
		// input wire(s) is not ready (doesn't have a value to work on yet)
		return
	}

	o.TargetWire.Signal = o.InputWire1.Signal | o.InputWire2.Signal
	o.TargetWire.simulate()
}

func (ls lshift) Formulate() {
	if !ls.InputWire.Simulated {
		// input wire is not ready (doesn't have a value to work on yet)
		return
	}

	ls.TargetWire.Signal = ls.InputWire.Signal << ls.ShiftAmount
	ls.TargetWire.simulate()
}

func (rs rshift) Formulate() {
	if !rs.InputWire.Simulated {
		// input wire is not ready (doesn't have a value to work on yet)
		return
	}

	rs.TargetWire.Signal = rs.InputWire.Signal >> rs.ShiftAmount
	rs.TargetWire.simulate()
}
