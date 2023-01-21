package polynomials






// Returns the real parts of a slice complex numbers
func getRealParts(cSlice []complex128) []float64{

	realParts := []float64{}

	for _, c := range cSlice{
		if imag(c) == 0 {
			realParts = append(realParts, real(c))
		}
	}

	return realParts
}