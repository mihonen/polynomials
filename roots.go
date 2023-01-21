package polynomials

import (
	"errors"
	"gonum.org/v1/gonum/mat"
)



// Package has three solving methods
// =================================
// 1. Durand-Kerner method 
//    https://en.wikipedia.org/wiki/Durand–Kerner_method
//
// 2. Root isolation + Newtons method 
//    https://en.wikipedia.org/wiki/Sturm%27s_theorem#Root_isolation
//	  https://en.wikipedia.org/wiki/Newton's_method
//
// 3. A method based on computing the companion matrix of the polynomial and solving its eigenvalues
// 	  https://en.wikipedia.org/wiki/Eigenvalue_algorithm#Algorithms
// 	  https://en.wikipedia.org/wiki/Companion_matrix
//
// The third method is usually the most robust


type SolvingMethod int

const (
    DurandKerner SolvingMethod = iota
    BisectionNewton
    Eigenvalue 
)



func (poly *Polynomial) PositiveRoots() ([]float64, error) {
	pRoots := []float64{}
	roots, err := poly.RealRoots()
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



func (poly *Polynomial) RealRoots() ([]float64, error){

	realRoots := []float64{}


	if poly.Degree() == 2 {
		complexRoots := poly.QuadraticRoots()
		realRoots = getRealParts(complexRoots)
	} else {
		switch poly.solveMode {
		case DurandKerner:

			complexRoots, err := poly.ComplexRootsDurandKerner()
			if err != nil {
				return realRoots, err
			}
			realRoots = getRealParts(complexRoots)

		case BisectionNewton:
			roots, err := poly.RootsBisectionNewton()
			if err != nil {
				return realRoots, err
			}
			realRoots = roots
			
		case Eigenvalue:
			complexRoots, err := poly.ComplexRootsEigenvalue();
			if err != nil {
				return realRoots, err
			}
			realRoots = getRealParts(complexRoots)
		}

	}

	return realRoots, nil
}


func (poly *Polynomial) ComplexRoots() ([]complex128, error){

	if poly.Degree() == 2{
		return poly.QuadraticRoots(), nil
	} else {
		switch poly.solveMode {
		case DurandKerner:
			return poly.ComplexRootsDurandKerner()

		case BisectionNewton:
			return []complex128{}, errors.New("BisectionNewton solve mode cannot solve complex roots. Change to either DurandKerner or Eigenvalue method.")

		case Eigenvalue:
			return poly.ComplexRootsEigenvalue()
		}

		return []complex128{}, errors.New("Invalid solve mode")
	}

}


func (poly *Polynomial) ComplexRootsDurandKerner() ([]complex128, error){
	roots, err := poly.DurandKernerRoots()
	if err != nil {
		return []complex128{}, err
	}

	for idx, root := range roots {
		roots[idx] = RoundC(root)
	}

	return roots, nil
}

func (poly *Polynomial) ComplexRootsEigenvalue() ([]complex128, error){
	poly.MakeMonic()
	companionMatrix, err := poly.CompanionMatrix()

	if err != nil {
		return []complex128{}, err
	}

	var eig mat.Eigen
	ok := eig.Factorize(companionMatrix, mat.EigenNone)
	if !ok {
		return []complex128{}, errors.New("Eigendecomposition failed")
	}


	roots := eig.Values(nil)
	for idx, root := range roots {
		roots[idx] = RoundC(root)
	}

	return roots, nil
}

func (poly *Polynomial) RootsBisectionNewton() ([]float64, error){
	lowerBound, upperBound := poly.RootBounds()
	roots, err := poly.RootsWithin(lowerBound, upperBound)

	if err != nil {
		return []float64{}, err
	}

	for idx, root := range roots {
		roots[idx] = Round(root)
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
		// log.Printf("%d ROOTS IN [%f, %f]", nRoots, a, b)
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









