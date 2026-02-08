// shapes/main.go
package main

import (
    "fmt"
)

func main() {
    // Create instances of each shape
    rect := Rectangle{Width: 5.0, Height: 10.0}
    circle := Circle{Radius: 7.0}
    triangle := Triangle{Base: 4.0, Height: 6.0}
    
    fmt.Println("=== Initial Shapes ===")
    fmt.Printf("Rectangle: Width=%.1f, Height=%.1f\n", rect.Width, rect.Height)
    fmt.Printf("  Area: %.2f, Perimeter: %.2f\n", rect.Area(), rect.Perimeter())
    
    fmt.Printf("\nCircle: Radius=%.1f\n", circle.Radius)
    fmt.Printf("  Area: %.2f, Perimeter: %.2f\n", circle.Area(), circle.Perimeter())
    
    fmt.Printf("\nTriangle: Base=%.1f, Height=%.1f\n", triangle.Base, triangle.Height)
    fmt.Printf("  Area: %.2f, Perimeter: %.2f\n", triangle.Area(), triangle.Perimeter())
    
    // Scale shapes
    fmt.Println("\n=== Scaling Shapes ===")
    rect.Scale(2.0)
    circle.Scale(1.5)
    triangle.Scale(3.0)
    
    fmt.Printf("Rectangle after scaling by 2: Width=%.1f, Height=%.1f\n", rect.Width, rect.Height)
    fmt.Printf("  Area: %.2f, Perimeter: %.2f\n", rect.Area(), rect.Perimeter())
    
    fmt.Printf("\nCircle after scaling by 1.5: Radius=%.1f\n", circle.Radius)
    fmt.Printf("  Area: %.2f, Perimeter: %.2f\n", circle.Area(), circle.Perimeter())
    
    fmt.Printf("\nTriangle after scaling by 3: Base=%.1f, Height=%.1f\n", triangle.Base, triangle.Height)
    fmt.Printf("  Area: %.2f, Perimeter: %.2f\n", triangle.Area(), triangle.Perimeter())
}