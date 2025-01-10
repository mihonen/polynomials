package polynomials

import (
	"fmt"
	"math"
	"strings"
)

// A Polynomial is represented as a slice of coefficients ordered increasingly by degree.
// eg. coeffs[0] * x^4 + coeffs[1] * x^3 coeffs[2] * x^2 ...
//

type Polynomial struct {
	coeffs     []float64
	sturmChain []*Polynomial
	SolveMode  SolvingMethod
}

// CreatePolynomial returns a new Polynomial
func CreatePolynomial(coefficients ...float64) *Polynomial {
	var newPolynomial Polynomial

	stripped := append([]float64{}, coefficients...)

	// Strip leading zeros
	for _, coeff := range coefficients {
		if coeff == 0.0 {
			stripped = stripped[1:]
		} else if math.IsNaN(coeff) {
			panic("Cannot create polynomial with NaN coefficient!")
		} else {
			break
		}
	}

	newPolynomial.coeffs = append([]float64{}, stripped...)
	//newPolynomial.RoundCoeffs()

	newPolynomial.SolveMode = DefaultSolvingMethod
	return &newPolynomial
}

func (poly *Polynomial) RoundCoeffs() {
	for idx, coeff := range poly.coeffs {
		poly.coeffs[idx] = Round(coeff)
	}
}

// Creates simple power polynomial, eg. x^3
func CreatePower(power int) *Polynomial {
	coeffs := []float64{}
	coeffs = append(coeffs, 1.0)

	for i := power; i > 0; i-- {
		coeffs = append(coeffs, 0.0)
	}

	return CreatePolynomial(coeffs...)
}

func (poly *Polynomial) Degree() int {
	// Coefficients should be maintained in such a way that allow the
	// number of coefficients to be one less than the degree of the polynomial.
	deg := len(poly.coeffs) - 1
	if deg < 0{
		return 0
	}
	return deg
}

func (poly *Polynomial) MakeMonic() {
	// Divides the polynomial with the leading coefficient to make the polynomial monic
	l := poly.LeadingCoeff()

	for idx, coeff := range poly.coeffs {
		poly.coeffs[idx] = coeff / l
	}
}

func (poly *Polynomial) IsMonic() bool {
	n := len(poly.coeffs)
	if n <= 1 {
		return false
	}

	return poly.coeffs[0] == 1.0
}

// At returns the value of the polynomial evaluated at x.
func (poly *Polynomial) At(x float64) float64 {
	// Implement Horner's Method
	n := len(poly.coeffs)
	if n == 0 {
		return 0
	}

	out := poly.coeffs[0]

	for i := 1; i < n; i++ {
		out = out*x + poly.coeffs[i]
	}

	return Round(out)
}

// AtComplex returns the value of the polynomial evaluated at z
func (poly *Polynomial) AtComplex(z complex128) complex128 {
	// Implement Horner's Method for complex input z
	t := complex(0, 0)
	if len(poly.coeffs) == 0 {
		return t
	}

	for _, c := range poly.coeffs {
		t = t*z + complex(c, 0)
	}

	return RoundC(t)
}

func (poly *Polynomial) IsZero() bool {
	return poly.Degree() == 0 && poly.coeffs[0] == 0.0
}

func (poly *Polynomial) computeSturmChain() {
	if poly.IsZero() {
		return
	}
	var sturmChain []*Polynomial
	var rem *Polynomial
	var tmp Polynomial

	sturmChain = append(sturmChain, poly)

	deriv := poly.Derivative()
	sturmChain = append(sturmChain, deriv)

	for i := 1; i < poly.Degree(); i++ {
		if sturmChain[i].Degree() == 0 {
			break
		}

		tmp = *sturmChain[i-1]

		_, rem = tmp.EuclideanDiv(sturmChain[i])

		sturmChain = append(sturmChain, rem.ScalarMult(-1))
	}

	poly.sturmChain = sturmChain
}

func (poly *Polynomial) LeadingCoeff() float64 {
	return poly.coeffs[0]
}

func (poly *Polynomial) Coeffs() []float64 {
	return poly.coeffs[:]
}

// EuclideanDiv aka. Polynomial Long Division
// divides the polynomial by another polynomial and returns the quotient and the remainder
//
// https://en.wikipedia.org/wiki/Polynomial_greatest_common_divisor#Euclidean_division
// https://rosettacode.org/wiki/Polynomial_long_division

func (poly1 *Polynomial) EuclideanDiv(poly2 *Polynomial) (*Polynomial, *Polynomial) {
	if poly1 == nil || poly2 == nil {
		panic("received nil *Polynomial")
	}

	if poly2.IsZero() {
		panic("EuclideanDiv division by zero")
	}

	quotDegree := poly1.Degree() - poly2.Degree()
	quotCoeffs := make([]float64, quotDegree+1)
	var d *Polynomial
	var shift int
	var factor float64

	r := poly1

	for r.Degree() >= poly2.Degree() {
		shift = r.Degree() - poly2.Degree()
		d = poly2.ShiftRight(shift)

		factor = r.LeadingCoeff() / d.LeadingCoeff()
		quotCoeffs[quotDegree-shift] = factor
		d = d.ScalarMult(factor)
		r = r.Sub(d)
	}

	quotient := CreatePolynomial(quotCoeffs...)
	return quotient, r
}

func (poly *Polynomial) ShiftRight(offset int) *Polynomial {
	if offset < 0 {
		panic("invalid offset")
	}
	shiftedCoeffs := make([]float64, len(poly.coeffs)+offset)
	copy(shiftedCoeffs, poly.coeffs)
	poly = CreatePolynomial(shiftedCoeffs...)
	return poly
}

// Subdivision of polynomials, returns result as a new polynomial
func (poly1 *Polynomial) Sub(poly2 *Polynomial) *Polynomial {
	var maxNumCoeffs int
	coeffs1 := poly1.coeffs
	coeffs2 := poly2.coeffs

	// Pad "shorter" polynomial with 0s.
	if len(coeffs1) > len(coeffs2) {
		maxNumCoeffs = len(coeffs1)
		for len(coeffs2) < maxNumCoeffs {
			coeffs2 = append(coeffs2, 0.0)
		}

	} else if len(coeffs1) < len(coeffs2) {
		maxNumCoeffs = len(coeffs2)
		for len(coeffs1) < maxNumCoeffs {
			coeffs1 = append(coeffs1, 0.0)
		}
	} else {
		maxNumCoeffs = len(coeffs1)
	}

	// Subtract coefficients with matching degrees.
	diffCoeffs := make([]float64, maxNumCoeffs)

	for i := 0; i < maxNumCoeffs; i++ {
		diffCoeffs[i] = coeffs1[i] - coeffs2[i]
	}

	newPoly := CreatePolynomial(diffCoeffs...)
	return newPoly
}

func (poly1 *Polynomial) Mult(poly2 *Polynomial) *Polynomial {
	prodCoeffs := make([]float64, poly1.Degree()+poly2.Degree()+1)

	for i := 0; i < len(poly1.coeffs); i++ {
		for j := 0; j < len(poly2.coeffs); j++ {
			prodCoeffs[i+j] += poly1.coeffs[i] * poly2.coeffs[j]
		}
	}

	prod := CreatePolynomial(prodCoeffs...)
	return prod
}

func (poly *Polynomial) ScalarMult(s float64) *Polynomial {

	coeffs := make([]float64, len(poly.coeffs))

	for i := 0; i < len(poly.coeffs); i++ {
		coeffs[i] = poly.coeffs[i] * s
	}

	newPoly := CreatePolynomial(coeffs...)

	return newPoly
}

func (poly1 *Polynomial) Add(poly2 *Polynomial) *Polynomial {
	coeffs1 := poly1.coeffs
	coeffs2 := poly2.coeffs

	// determine "longer" and "shorter" coeffs

	var longer, shorter *[]float64
	if len(coeffs1) > len(coeffs2) {
		longer = &coeffs1
		shorter = &coeffs2
	} else {
		longer = &coeffs2
		shorter = &coeffs1
	}

	// pad "shorter" coeff with zeros

	delta := len(*longer) - len(*shorter)
	for ; delta > 0; delta-- {
		*shorter = append([]float64{0}, *shorter...)
	}

	// add corresponding coeffs

	coeffsSum := make([]float64, len(*longer))
	for i := range coeffsSum {
		coeffsSum[i] = coeffs1[i] + coeffs2[i]
	}

	return CreatePolynomial(coeffsSum...)
}

// String returns a string representation of the polynomial
func (poly *Polynomial) String() string {
	lc := len(poly.coeffs)
	if lc == 0 {
		return "0"
	}

	if lc == 1 {
		if poly.coeffs[0] > 0 {
			return fmt.Sprintf("%0.3f", poly.coeffs[0])
		} else {
			return fmt.Sprintf("- %0.3f", -poly.coeffs[0])
		}
	}

	var s = strings.Builder{}
	for i := 0; i < lc; i++ {
		if poly.coeffs[i] == 0 {
			continue
		}

		coeff := poly.coeffs[i]
		sign := " + "
		if poly.coeffs[i] < 0 {
			coeff = -coeff
			sign = " - "
		}

		if i == 0 {
			if poly.coeffs[i] < 0 {
				s.WriteString(sign[1:])
			}
			if lc > 2 {
				s.WriteString(fmt.Sprintf("%0.3fx^%d", coeff, lc-1))
				continue
			} else {
				s.WriteString(fmt.Sprintf("%0.3fx", coeff))
				continue
			}
		} else if i == lc-1 {
			s.WriteString(sign)
			s.WriteString(fmt.Sprintf("%0.3f", coeff))
		} else if i == lc-2 {
			s.WriteString(sign)
			s.WriteString(fmt.Sprintf("%0.3fx", coeff))
		} else {
			s.WriteString(sign)
			s.WriteString(fmt.Sprintf("%0.3fx^%d", coeff, lc-i-1))
		}
	}
	return s.String()
}
