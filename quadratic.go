package polynomials

import (
	"math"
)




func (poly *Polynomial) QuadraticRoots() ([]complex128){
	if poly.Degree() != 2 { panic("cannot use quatratic formula on non-quatratic polynomial") }

	a := poly.coeffs[0]
	b := poly.coeffs[1]
	c := poly.coeffs[2]

	discriminant := b*b - 4.0*a*c
	if discriminant == 0 {

		realPart := (-b) / (2.0*a)
		root := complex(realPart, 0)

		return[]complex128{root}

	} else if discriminant > 0 {

		real1 := (-b + math.Sqrt(discriminant) ) / (2.0*a)
	    real2 := (-b - math.Sqrt(discriminant) ) / (2.0*a)

	    root1 := complex(real1, 0)
	    root2 := complex(real2, 0)

	    return []complex128{root1, root2}

	} else { // complex roots
		realPart := -b / (2.0*a)
		imgPart  := math.Sqrt(-discriminant) / (2.0*a)
		root1 := complex(realPart, +imgPart)
		root2 := complex(realPart, -imgPart)


		return []complex128{root1, root2}
	}


}

