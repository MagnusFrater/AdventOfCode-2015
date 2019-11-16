package circuit

import "fmt"

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
		fmt.Printf("%v: already simulated\n", w.Name)
		return
	}
	w.Simulated = true

	fmt.Printf("\n%v: simulating connections...\n", w.Name)

	for _, connection := range w.Connections {
		(*connection).Formulate()
	}

	fmt.Printf("%v: finished simulating connections\n\n", w.Name)
}
