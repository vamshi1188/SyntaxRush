package main

import (
	"fmt"
	"math"
)

// Calculator represents a simple calculator
type Calculator struct {
	memory float64
}

// NewCalculator creates a new calculator instance
func NewCalculator() *Calculator {
	return &Calculator{memory: 0}
}

// Add performs addition
func (c *Calculator) Add(a, b float64) float64 {
	result := a + b
	c.memory = result
	return result
}

// Multiply performs multiplication
func (c *Calculator) Multiply(a, b float64) float64 {
	result := a * b
	c.memory = result
	return result
}

// Divide performs division with error handling
func (c *Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	result := a / b
	c.memory = result
	return result, nil
}

// SquareRoot calculates square root
func (c *Calculator) SquareRoot(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("square root of negative number")
	}
	result := math.Sqrt(a)
	c.memory = result
	return result, nil
}

// GetMemory returns the stored memory value
func (c *Calculator) GetMemory() float64 {
	return c.memory
}

// main demonstrates the calculator usage
func main() {
	calc := NewCalculator()
	
	fmt.Println("=== Calculator Demo ===")
	
	// Basic operations
	sum := calc.Add(10, 5)
	fmt.Printf("10 + 5 = %.2f\n", sum)
	
	product := calc.Multiply(4, 7)
	fmt.Printf("4 * 7 = %.2f\n", product)
	
	// Division with error handling
	if quotient, err := calc.Divide(15, 3); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("15 / 3 = %.2f\n", quotient)
	}
	
	// Square root with error handling
	if sqrt, err := calc.SquareRoot(16); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("√16 = %.2f\n", sqrt)
	}
	
	// Memory operations
	fmt.Printf("Memory: %.2f\n", calc.GetMemory())
	
	fmt.Println("Calculator demo completed!")
}
