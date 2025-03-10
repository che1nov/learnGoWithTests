package main

import (
	"math"
)

type Shape interface {
	Area() float64
}

type Restangle struct {
	Width  float64
	Height float64
}

func (r Restangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

func Perimeter(restangle Restangle) float64 {
	return 2 * (restangle.Width + restangle.Height)
}

func Area(restangle Restangle) float64 {
	return restangle.Width * restangle.Height
}
