package polynomials

import (
	"math"
)




func (poly *Polynomial) QuadraticRoots() ([]float64){
	if poly.Degree() != 2 { panic("cannot use quatratic formula on non-quatratic polynomial") }

	a := poly.coeffs[0]
	b := poly.coeffs[1]
	c := poly.coeffs[2]

	discriminant := b*b - 4.0*a*c
	if discriminant == 0 {

		x := (-b) / (2.0*a)
		return[]float64{x}

	} else if discriminant > 0 {

		x1 := (-b + math.Sqrt(discriminant) ) / (2.0*a)
	    x2 := (-b - math.Sqrt(discriminant) ) / (2.0*a)
	    return []float64{x1, x2}

	} else {
		return []float64{}
	}


}

