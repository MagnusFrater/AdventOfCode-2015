package day09

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type graph map[string]*vertex

func (g graph) createVertex(location string) {
	if _, prs := g[location]; !prs {
		g[location] = &vertex{
			Name: location,
		}
	}
}

func (g graph) addEdges(edges []string) {
	for _, edge := range edges {
		var split = strings.Fields(edge)
		g.addEdge(split[0], split[2], split[4])
	}
}

func (g graph) addEdge(vertex1, vertex2, weight string) {
	g.createVertex(vertex1)
	g.createVertex(vertex2)

	weightInt, err := strconv.Atoi(weight)
	if err != nil {
		panic(err)
	}

	g[vertex1].addEdge(g[vertex2], weightInt)
	g[vertex2].addEdge(g[vertex1], weightInt)
}

// reset sets all vertices to unvisited
func (g graph) reset() {
	for _, vertex := range g {
		vertex.Visited = false
	}
}

func (g graph) getShortestDistance() int {
	var shortestDistance = math.MaxInt64

	for _, city := range g {
		g.reset()

		var distance = city.getShortestDistance()
		if distance < shortestDistance {
			shortestDistance = distance
		}
	}

	return shortestDistance
}

func (g graph) getLongestDistance() int {
	var longestDistance = math.MinInt64

	for _, city := range g {
		g.reset()

		var distance = city.getLongestDistance()
		if distance > longestDistance {
			longestDistance = distance
		}
	}

	return longestDistance
}

func (g graph) print() {
	for name, vertex := range g {
		fmt.Println(name)

		for _, edge := range vertex.Edges {
			fmt.Printf("  -- (%v) --> %v\n", edge.Weight, edge.Vertex.Name)
		}

		fmt.Println()
	}
}

func (g graph) printClosestNeighbours() {
	for name, city := range g {
		var closestNeighbour = city.getClosestUnvisitedNeighbour()
		fmt.Printf("%v -- (%v) --> %v\n", name, closestNeighbour.Weight, closestNeighbour.Vertex.Name)
	}
	fmt.Println()
}

func (v *vertex) addEdge(vertex *vertex, weight int) {
	v.Edges = append(v.Edges, &edge{vertex, weight})
}
