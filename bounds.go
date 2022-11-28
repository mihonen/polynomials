package polynomials

import (
	"math"
)

// https://en.wikipedia.org/wiki/Geometrical_properties_of_polynomial_roots#Bounds_on_(complex)_polynomial_roots



// Cauchy's bound

func (poly *Polynomial) RootBounds() (float64, float64){

	n := len(poly.coeffs)
	var a float64
	var maxA float64 = 0

	for i := 0; i < n; i++ {
		a = math.Abs(poly.coeffs[i] / poly.coeffs[0])
		if a > maxA {
			maxA = a
		}
	}

	upperBound := 1.0 + maxA
	lowerBound := -upperBound

	return lowerBound, upperBound
}






