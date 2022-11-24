package polynomials

import (
	"errors"
)


func (poly *Polynomial) PositiveRoots() ([]float64, error){
	return poly.Roots()
}

func (poly *Polynomial) Roots() ([]float64, error){
	if poly.Degree() == 2 {
		return poly.QuadraticRoots(), nil
	} else {
		return []float64{}, errors.New("Roots() not implemented for real roots yet!")
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
	return []float64{}, nil
}


func (poly *Polynomial) NumericalRoots() ([]float64, error){
	return []float64{}, nil
}

