package polynomials

import (

)

// A Polynomial is represented as a slice of coefficients ordered increasingly by degree.

type Polynomial struct {
	coeffs []float64
}


// CreatePolynomial returns a new Polynomial
func CreatePolynomial(coefficients ...float64) (*Polynomial) {
	var newPolynomial Polynomial
	newPolynomial.coeffs = append([]float64{}, coefficients...)
	return &newPolynomial
}


func (poly *Polynomial) Degree() int {
	// Coefficients should be maintained in such a way that allow the
	// number of coefficients to be one less than the degree of the polynomial.
	return len(poly.coeffs) - 1
}


// At returns the value of the polynomial evaluated at x.
func (poly *Polynomial) At(x float64) float64 {
	// Implement Horner's Method
	n := len(poly.coeffs)
	out := poly.coeffs[0]

	for i := 1; i < n; i++ {
		out = out * x + poly.coeffs[i]
	}    

	return Round(out)
}


// AtComplex returns the value of the polynomial evaluated at z
func (poly *Polynomial) AtComplex(z complex128) complex128 {
	// Implement Horner's Method for complex input z
	t := complex(0, 0)
	for _, c := range poly.coeffs {
		t = t * z + complex(c, 0)
	}

    return RoundC(t)
}


