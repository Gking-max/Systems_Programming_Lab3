// shapes/shapes.go
package main

import "math"

// Rectangle struct
type Rectangle struct {
    Width  float64
    Height float64
}

// Circle struct
type Circle struct {
    Radius float64
}

// Triangle struct
type Triangle struct {
    Base   float64
    Height float64
}

// Area methods with value receivers
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
    return 0.5 * t.Base * t.Height
}

// Perimeter methods with value receivers
func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

func (t Triangle) Perimeter() float64 {
    // Assuming equilateral triangle
    return 3 * t.Base
}

// Scale methods with pointer receivers
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

func (c *Circle) Scale(factor float64) {
    c.Radius *= factor
}

func (t *Triangle) Scale(factor float64) {
    t.Base *= factor
    t.Height *= factor
}