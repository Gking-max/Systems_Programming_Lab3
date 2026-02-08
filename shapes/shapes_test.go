// shapes/shapes_test.go
package main

import (
    "math"
    "testing"
)

func TestRectangleArea(t *testing.T) {
    rect := Rectangle{Width: 5.0, Height: 10.0}
    expected := 50.0
    got := rect.Area()
    if math.Abs(got-expected) > 0.0001 {
        t.Errorf("Rectangle Area() = %.2f, want %.2f", got, expected)
    }
}

func TestRectanglePerimeter(t *testing.T) {
    rect := Rectangle{Width: 5.0, Height: 10.0}
    expected := 30.0
    got := rect.Perimeter()
    if math.Abs(got-expected) > 0.0001 {
        t.Errorf("Rectangle Perimeter() = %.2f, want %.2f", got, expected)
    }
}

func TestRectangleScale(t *testing.T) {
    rect := &Rectangle{Width: 5.0, Height: 10.0}
    rect.Scale(2.0)
    if rect.Width != 10.0 || rect.Height != 20.0 {
        t.Errorf("Rectangle Scale() failed: got Width=%.1f, Height=%.1f, want Width=10.0, Height=20.0", 
            rect.Width, rect.Height)
    }
}

func TestCircleArea(t *testing.T) {
    circle := Circle{Radius: 7.0}
    expected := math.Pi * 49.0
    got := circle.Area()
    if math.Abs(got-expected) > 0.0001 {
        t.Errorf("Circle Area() = %.2f, want %.2f", got, expected)
    }
}

func TestCirclePerimeter(t *testing.T) {
    circle := Circle{Radius: 7.0}
    expected := 2 * math.Pi * 7.0
    got := circle.Perimeter()
    if math.Abs(got-expected) > 0.0001 {
        t.Errorf("Circle Perimeter() = %.2f, want %.2f", got, expected)
    }
}

func TestCircleScale(t *testing.T) {
    circle := &Circle{Radius: 7.0}
    circle.Scale(2.0)
    if circle.Radius != 14.0 {
        t.Errorf("Circle Scale() failed: got Radius=%.1f, want 14.0", circle.Radius)
    }
}

func TestTriangleArea(t *testing.T) {
    triangle := Triangle{Base: 4.0, Height: 6.0}
    expected := 12.0
    got := triangle.Area()
    if math.Abs(got-expected) > 0.0001 {
        t.Errorf("Triangle Area() = %.2f, want %.2f", got, expected)
    }
}

func TestTrianglePerimeter(t *testing.T) {
    triangle := Triangle{Base: 4.0, Height: 6.0}
    expected := 12.0 // 3 * base for equilateral triangle
    got := triangle.Perimeter()
    if math.Abs(got-expected) > 0.0001 {
        t.Errorf("Triangle Perimeter() = %.2f, want %.2f", got, expected)
    }
}

func TestTriangleScale(t *testing.T) {
    triangle := &Triangle{Base: 4.0, Height: 6.0}
    triangle.Scale(3.0)
    if triangle.Base != 12.0 || triangle.Height != 18.0 {
        t.Errorf("Triangle Scale() failed: got Base=%.1f, Height=%.1f, want Base=12.0, Height=18.0",
            triangle.Base, triangle.Height)
    }
}