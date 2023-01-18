package main

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
	A      float64
	B      float64
	C      float64
}

type Shape interface {
	Area() float64
	Parameter() float64
}
