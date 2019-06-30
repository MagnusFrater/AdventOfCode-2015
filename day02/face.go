package day02

// Face is a flat (planar) surface that forms part of the boundary of a solid object.
type Face interface {
	Area() float64
	Perimeter() float64
}

// Rectangle is a quadrilateral with four right angles
type Rectangle struct {
	x, y float64
}

// Area returns the area of the Rectangle.
func (r Rectangle) Area() float64 {
	return r.x * r.y
}

// Perimeter returns the Perimeter of the Rectangle.
func (r Rectangle) Perimeter() float64 {
	return 2*r.x + 2*r.y
}
