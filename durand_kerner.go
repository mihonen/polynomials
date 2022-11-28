package polynomials

import (
	//"log"
	"math"
	"math/rand"
	//"errors"
)


func (poly *Polynomial) EhrlichRadius() float64 {
	if poly.Degree() == 0 { return 1.0 }

    n := poly.Degree()
    a_0 := poly.coeffs[0]
    a_n := poly.coeffs[n]

    abs := math.Abs(a_0 / a_n)

    r   := math.Pow(abs, 1.0 / float64(n))

    return r
}

func (poly *Polynomial) Bound() float64 {
	

    n := len(poly.coeffs) - 1
    b := 0.0

	reversed := make([]float64, len(poly.coeffs))
	copy(reversed, poly.coeffs)
	Reverse(reversed)

    for i := 0; i < n; i++ {
        b += math.Abs(reversed[i] / reversed[i + 1])   
    }

    return b
}


func (poly *Polynomial) DurandKernerRoots() ([]complex128, error){
	n := poly.Degree()
	roots    := make([]complex128, n )
	// rootsNew := make([]complex128, n)
	// theta  := 2.0 * math.Pi / float64(n)
	bnd    := poly.Bound()

	for k := 0; k < n; k++ {
		r := bnd * rand.Float64()
		theta := 2.0 * math.Pi * rand.Float64()
		roots[k] = complex(r * math.Cos(theta), r * math.Sin(theta))
	}

	max_delta := 1.0

	for i := 0; i < durandKernerMaxIter; i++ {
		max_delta = 0
		for k := 0; k < n; k++ {
			// deno := complex(1.0, 0.0)
			delta := poly.AtComplex(roots[k])
			for j := 0; j < n; j++ {

			    if j != k {
			        delta /= (roots[k] - roots[j])
			    }
			}
			roots[k] -= delta

			max_delta += ((real(delta) * real(delta)) +
			              (imag(delta) * imag(delta))) / float64(n)
		}
		               
		max_delta = max_delta * max_delta
		if max_delta < eps {
			break
		}           
	}
	
	return roots, nil
}



