package polynomials

import (
	"errors"
)


func (poly *Polynomial) PositiveRoots() ([]float64, error){
	pRoots := []float64{}
	roots, err := poly.Roots()
	if err != nil {
		return []float64{}, err
	}

	for _, root := range roots {
		if root > 0 {
			pRoots = append(pRoots, root)
		}
	}

	return pRoots, nil
}

func (poly *Polynomial) Roots() ([]float64, error){
	if poly.Degree() == 2 {
		return poly.QuadraticRoots(), nil
	} else {
		rRoots := []float64{}
		cRoots, err := poly.ComplexRoots()
		if err != nil {
			return []float64{}, err
		}

		for _, cRoot := range cRoots{
			if imag(cRoot) == 0 {
				rRoots = append(rRoots, real(cRoot))
			}
		}
		return rRoots, nil
	}
}


func (poly *Polynomial) ComplexRoots() ([]complex128, error){
	roots, err := poly.DurandKernerRoots()
	if err != nil {
		return []complex128{}, err
	}

	for idx, root := range roots {
		roots[idx] = RoundC(root)
	}

	return roots, nil
}


func (poly *Polynomial) RootsWithin(lowerBound complex128, upperBound complex128) ([]float64, error){
	return []float64{}, errors.New("RootsWithin() not implemented yet!")
}



