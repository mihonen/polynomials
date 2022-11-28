package polynomials






type Interval struct {
	A float64
	B float64
}



func (i *Interval) Mid() float64 {
	return (i.A + i.B) / 2.0
}