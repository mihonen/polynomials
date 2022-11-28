package polynomials


import (
	"math"
)

// Because of floating point errors, 
// results should be rounded to some number of decimal places 
// to perform precise arithmetics 
// ==============================
// FURTHER INFO: 
// https://en.wikipedia.org/wiki/Floating-point_arithmetic
// https://www.geeksforgeeks.org/floating-point-error-in-python/


func Round(value float64) float64 {
	n := math.Pow(10.0, float64(roundingDecimalPlaces))
	rounded := float64(math.Round(n * value) / n)
	return rounded
}

// Complex Round
func RoundC(z complex128) complex128{
	n := math.Pow(10.0, float64(roundingDecimalPlaces))
	roundedA := float64(math.Round(n * real(z)) / n)
	roundedB := float64(math.Round(n * imag(z)) / n)

	return complex(roundedA, roundedB)
}