package circuit

// wire is a vertex in the graph.
type wire struct {
	Name        string
	Signal      uint16
	Connections []*formula
	Simulated   bool
}

func (w *wire) addConnection(f formula) {
	w.Connections = append(w.Connections, &f)
}

func (w *wire) simulate() {
	if w.Simulated {
		return
	}
	w.Simulated = true

	for _, connection := range w.Connections {
		(*connection).Formulate()
	}
}
