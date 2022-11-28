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
		lowerBound, upperBound := poly.RootBounds()
		return poly.RootsWithin(lowerBound, upperBound)
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


func (poly *Polynomial) RootsWithin(lowerBound float64, upperBound float64) ([]float64, error){

	if poly.IsZero() {
		return nil, errors.New("infinitely many solutions")
	}

	roots := []float64{}
	// Check if lowerBound is a root
	if poly.At(lowerBound) == 0.0 {
		roots = append(roots, lowerBound)
	}


	isolationIntervals := poly.findIsolationIntervals(lowerBound, upperBound)

	for _, isolationInterval := range isolationIntervals {
		root, err := poly.NewtonMethod(isolationInterval.Mid())
		if err != nil {
			return roots, err
		}

		roots = append(roots, root)
	}


	return roots, nil

}

	
// Returns an array of intervals, where each intervals holds one root

func (poly *Polynomial) findIsolationIntervals(a float64, b float64) ([]Interval) {
	isolationIntervals := []Interval{}

	nRoots := poly.countRootsWithin(a, b)

	if nRoots > 1 {
		// Divide interval further into two intervals
		mp := (a + b) / 2.0
		intervals1 := poly.findIsolationIntervals(a, mp)
		intervals2 := poly.findIsolationIntervals(mp, b)


		isolationIntervals = append(isolationIntervals, intervals1...)
		isolationIntervals = append(isolationIntervals, intervals2...)

	} else if nRoots == 1 {
		interval := Interval{A: a, B: b}
		isolationIntervals = append(isolationIntervals, interval)
	} else { 
		return []Interval{}
	}

	return isolationIntervals
}




func (poly *Polynomial) countRootsWithin(a, b float64) int {
	if len(poly.sturmChain) == 0 {
		poly.computeSturmChain()
	}


	var seqA, seqB []float64

	for _, p := range poly.sturmChain {
		seqA = append(seqA, p.At(a))
		seqB = append(seqB, p.At(b))
	}
	return signVar(seqA) - signVar(seqB)
}



func signVar(s []float64) int {

	var filtered []float64
	for i := 0; i < len(s); i++ {
		if s[i] != 0.0 {
			filtered = append(filtered, s[i])
		}
	}

	// Count sign variations
	var count int
	for i := 0; i < len(filtered)-1; i++ {
		if filtered[i]*filtered[i+1] < 0 {
			count++
		}
	}

	return count
}









