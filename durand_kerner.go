package polynomials

import (
	"math"
	"errors"
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



func (poly *Polynomial) DurandKernerRoots() ([]complex128, error){
	n := poly.Degree()

	roots    := make([]complex128, n )
	rootsNew := make([]complex128, n)

	radius := poly.EhrlichRadius()
	theta  := 2.0 * math.Pi / float64(n)
	offset := theta / (float64(n) + 1.0)

	retry  := true

	for retry{
		retry = false
		// set initial roots as uniformly distributed points within ehrlich circle
		for k := 0; k < n; k++ {
			roots[k] = complex(radius * math.Cos(float64(k) * theta + offset), radius * math.Sin(float64(k) * theta + offset))
		}

		itCtr := 0
		flag := true


		for flag {
			flag = false
			for k := 0; k < n; k++ {
				temp := complex(1.0, 0.0)
				for j := 0; j < n; j++ {
				    if j != k {
				        temp *= roots[k] - roots[j]
				    }
				}

				rootsNew[k] = roots[k] - poly.AtComplex(roots[k]) / temp

				if math.Abs(real(roots[k]) - real(rootsNew[k])) > eps{

					flag = true
				}

				if math.IsNaN(real(rootsNew[k])) || math.IsNaN(imag(rootsNew[k])){
					flag = false
					retry = true
					break
				}


			}


			copy(roots, rootsNew)
			itCtr += 1
			if itCtr > durandKernerMaxIter {
				return []complex128{}, errors.New("DurandKerner method failed to converge before max number of iterations was exceeded!")
			}
		}

	}

	
	return roots, nil
}



