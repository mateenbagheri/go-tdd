package main

import "math"

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Parameter() float64 {
	return (r.Width + r.Height) * 2
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Parameter() float64 {
	return math.Pi * c.Radius * 2
}
