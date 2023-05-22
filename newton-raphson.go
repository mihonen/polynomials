package polynomials



import (
	"math"
	"errors"
)

// Newton's Method Implementation
// https://en.wikipedia.org/wiki/Newton%27s_method
func (poly *Polynomial) NewtonMethod(guess float64) (float64, error) {

	deriv := poly.Derivative()
	root  := guess
	prev  := root

	var derivAtRoot float64
	for i := 0; i < MaxNewtonIterations; i++ {
		derivAtRoot = deriv.At(root)
		// In the case that the derivative evaluates to zero, return the current guess.
		if derivAtRoot == 0.0 {
			return root, nil
		}
		prev = root
		root -= poly.At(root) / derivAtRoot
		if math.Abs(prev - root) < EpsNewton {
			return root, nil
		}
	}

	return root, errors.New("NewtonRaphson didn't converge before max number of iteration was reached! Result may be incorrect")
}