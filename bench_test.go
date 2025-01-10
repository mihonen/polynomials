package polynomials



import (
	// "fmt"
	"testing"
)

func BenchmarkDurandKerner(t *testing.B){

		poly := CreatePolynomial(1.13, -5.0, 12.0, -2.8, 3.213)


		poly.SolveMode = DurandKerner
		// start := time.Now()
		_, err := poly.RealRoots()
		if err != nil {
			t.Fatalf(`%v`, err)
		}

		// fmt.Println("DURAND KERNER TOOK: ", time.Since(start))


}

func BenchmarkEigenvalue(t *testing.B){

		poly := CreatePolynomial(1.13, -5.0, 12.0, -2.8, 3.213)


		// fmt.Println("DURAND KERNER TOOK: ", time.Since(start))

		poly.SolveMode = Eigenvalue
		// start = time.Now()
		_, err := poly.RealRoots()
		if err != nil {
			t.Fatalf(`%v`, err)

		}

}


func BenchmarkNewton(t *testing.B){

		poly := CreatePolynomial(1.13, -5.0, 12.0, -2.8, 3.213)



		poly.SolveMode = BisectionNewton
		// start = time.Now()
		_, err := poly.RealRoots()
		if err != nil {
			t.Fatalf(`%v`, err)

		}

}


func BenchmarkQuadradic(t *testing.B){

		poly := CreatePolynomial(1.13, -5.0, 12.0)


		// start = time.Now()
		_, err := poly.RealRoots()
		if err != nil {
			t.Fatalf(`%v`, err)

		}

}






