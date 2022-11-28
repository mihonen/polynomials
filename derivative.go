package polynomials


func (poly *Polynomial) Derivative() *Polynomial {
	if poly.Degree() == 0 {
		deriv := CreatePolynomial(0)
		return deriv
	}

	nDerivativeCoeffs := len(poly.coeffs) - 1
	derivativeCoeffs := make([]float64, nDerivativeCoeffs)
	for i := 0; i < nDerivativeCoeffs; i++ {

		derivativeCoeffs[i] = poly.coeffs[i] * float64(nDerivativeCoeffs - i)
	}

	deriv := CreatePolynomial(derivativeCoeffs...)
	return deriv
}


