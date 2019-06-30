package day02

import (
	"sort"
	"strconv"
	"strings"
)

// Cuboid is a convex polyhedron bounded by six quadrilateral faces,
// whose polyhedral graph is the same as that of a cube.
type Cuboid interface {
	SurfaceArea() float64
	Volume() float64
	SmallestFace() Rectangle
}

// RectangularCuboid is a cuboid in which each of the faces is a rectangle.
type RectangularCuboid struct {
	l, w, h float64
}

func newRectangularCuboid(desc string) RectangularCuboid {
	split := strings.Split(desc, "x")

	l, _ := strconv.ParseFloat(split[0], 64)
	w, _ := strconv.ParseFloat(split[1], 64)
	h, _ := strconv.ParseFloat(split[2], 64)

	return RectangularCuboid{l, w, h}
}

// SurfaceArea returns the surface area of the RectangularCuboid.
func (rp RectangularCuboid) SurfaceArea() float64 {
	return 2 * (rp.l*rp.w + rp.w*rp.h + rp.h*rp.l)
}

// Volume returns the volume of the RectangularCuboid.
func (rp RectangularCuboid) Volume() float64 {
	return rp.l * rp.w * rp.h
}

// SmallestFace returns the smallest face of the RectangularCuboid.
func (rp RectangularCuboid) SmallestFace() Rectangle {
	n := []float64{rp.l, rp.w, rp.h}
	sort.Float64s(n)
	return Rectangle{n[0], n[1]}
}
