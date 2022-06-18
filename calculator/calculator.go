// Package calculator provides a library for simple
// calculations in Go.
package calculator

// Add takes two numbers and returns the result of adding
// them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers a and b, and
// returns the result of subtracting b from a.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers a and b, and
// returns the result of multiplying a and b.
func Multiply(a, b float64) float64 {
	return a * b
}

// Devide takes two numbers a and b, and
// returns the result of deviding a by b.
func Devide(a, b float64) (float64, error) {
	return a / b, nil
}
