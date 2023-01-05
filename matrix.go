package polynomials


import (
	"errors"
	"gonum.org/v1/gonum/mat"
)



// Computes the companion matrix of a monic polynomial
// REFER TO: https://en.wikipedia.org/wiki/Companion_matrix

func (poly *Polynomial) CompanionMatrix() (*mat.Dense, error) {
	if !poly.IsMonic() {
		return nil, errors.New("Polynomial is not monic. Cannot create companion matrix")
	}

	n := poly.Degree()

	matrix := mat.NewDense(n, n, nil)

	lastCol := []float64{}

	for i := n;  i > 0; i-- {
		lastCol = append(lastCol, -poly.coeffs[i])
	}

	matrix.SetCol(n-1, lastCol)

	for i := 0; i < n-1; i++ {
		col := make([]float64, n)
		col[i + 1] = 1
		matrix.SetCol(i, col)
	}

	return matrix, nil
}