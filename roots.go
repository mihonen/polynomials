package polynomials




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



func (poly *Polynomial) RootsWithin(lowerBound complex128, upperBound complex128) ([]float64, error){
	return []float64{}, nil
}


func (poly *Polynomial) NumericalRoots() ([]float64, error){
	return []float64{}, nil
}

