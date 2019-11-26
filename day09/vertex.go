package day09

import "math"

type (
	vertex struct {
		Name    string
		Visited bool
		Edges   []*edge
	}

	edge struct {
		Vertex *vertex
		Weight int
	}
)

func (v *vertex) getClosestUnvisitedNeighbour() *edge {
	var closest = &edge{
		Weight: math.MaxInt64,
	}

	for _, e := range v.Edges {
		if !e.Vertex.Visited && e.Weight < closest.Weight {
			closest = e
		}
	}

	if closest.Vertex == nil {
		return nil
	}

	return closest
}

func (v *vertex) getFurthestUnvisitedNeighbour() *edge {
	var furthest = &edge{
		Weight: math.MinInt64,
	}

	for _, e := range v.Edges {
		if !e.Vertex.Visited && e.Weight > furthest.Weight {
			furthest = e
		}
	}

	if furthest.Vertex == nil {
		return nil
	}

	return furthest
}

func (v *vertex) getShortestDistance() int {
	v.Visited = true

	var neighbour = v.getClosestUnvisitedNeighbour()
	if neighbour == nil {
		return 0
	}

	return neighbour.Weight + neighbour.Vertex.getShortestDistance()
}

func (v *vertex) getLongestDistance() int {
	v.Visited = true

	var neighbour = v.getFurthestUnvisitedNeighbour()
	if neighbour == nil {
		return 0
	}

	return neighbour.Weight + neighbour.Vertex.getLongestDistance()
}
