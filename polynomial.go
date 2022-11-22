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
	length := len(poly.coeffs)
	out := poly.coeffs[length-1]
	for i := length - 2; i >= 0; i-- {
		out = out*x + poly.coeffs[i]
	}
	return out
}

func (poly *Polynomial) PositiveRoots() ([]float64, error){
	return poly.Roots()
}

func (poly *Polynomial) Roots() ([]float64, error){
	if poly.Degree() == 2 {
		return poly.QuadraticRoots(), nil
	} else {
		return poly.NumericalRoots()
	}
}

func (poly *Polynomial) NumericalRoots() ([]float64, error){
	return []float64{}, nil
}

func (poly *Polynomial) Derivative() *Polynomial {
	if poly.Degree() == 0 {
		deriv := CreatePolynomial(0)
		return deriv
	}

	nDerivativeCoeffs := len(poly.coeffs) - 1
	derivativeCoeffs := make([]float64, nDerivativeCoeffs)
	for i := 0; i < nDerivativeCoeffs; i++ {
		derivativeCoeffs[i] = poly.coeffs[i+1] * float64(i+1)
	}

	deriv := CreatePolynomial(derivativeCoeffs...)
	return deriv
}