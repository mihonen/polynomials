package polynomials






import (
	"fmt"
	"time"
    "testing"
)





func TestBench(t *testing.T){
	fmt.Println("BENCHMARKS")
	poly := CreatePolynomial(1.13, -5.0, 12.0, -2.8, 3.213)


	poly.solveMode = DurandKerner
	start := time.Now()
	_, err := poly.Roots()
	if err != nil {
		t.Fatalf(`%v`, err)

	}
	fmt.Println("DURAND KERNER TOOK: ", time.Since(start))

	poly.solveMode = Eigenvalue
	start = time.Now()
	_, err = poly.Roots()
	if err != nil {
		t.Fatalf(`%v`, err)

	}
	fmt.Println("EIGENVALUE TOOK: ", time.Since(start))

	poly.solveMode = BisectionNewton
	start = time.Now()
	_, err = poly.Roots()
	if err != nil {
		t.Fatalf(`%v`, err)

	}
	fmt.Println("BisectionNewton TOOK: ", time.Since(start))


}