package polynomials


import (
	"fmt"
    "testing"
)


func TestPolynomialLongDivision(t *testing.T){
	N := CreatePolynomial(1.0, -5.0, 12.0, -2.8)
	D := CreatePolynomial(1.0, -3.0)

	q, _ := N.EuclideanDiv(D)

	if q.Degree() != 2 || q.coeffs[0] != 1.0 || q.coeffs[1] != -2.0 || q.coeffs[2] != 6.0 {
		t.Fatalf(`EuclideanDiv() returned wrong quotient: %v. True quotient: %f %f %f`, q.coeffs, 1.0, -2.0, 6.0)
		fmt.Println()
	}

	fmt.Println("EuclideanDiv .......... OK")

}

func TestShiftRight(t *testing.T){
	poly := CreatePolynomial(1.0, 2.0, 3.0)

	shifted := poly.ShiftRight(2)

	if shifted.Degree() != 4 || shifted.coeffs[0] != 1.0 {
		t.Fatalf(`ShiftRight() returned wrong solution: %v. True solution: %v`, shifted, []float64{1.0, 2.0, 3.0, 0.0, 0.0})
	}

	fmt.Println("ShiftRight ............ OK")

}


func TestStuckLoop(t *testing.T){
	// Test with high precision floats that they don't cause errors that block the package
	coeffs := []float64{-0.0005148053170874375, 0.01177362691607392, -0.10824061093058787, 0.07523007124191312, -0.4864048905537971}
	poly := CreatePolynomial(coeffs...)

	_, err := poly.Roots()
	if err != nil {
		t.Fatalf(`Roots() errored: %v`, err)
	}

	fmt.Println("Infinite Loop Test .... OK")
}

func TestNoSolution(t *testing.T){

	a := 1.0
	b := -6.0
	c := 15.0
	d := -18.0
	e := 10.0

	poly := CreatePolynomial(a, b, c, d, e)

	roots, err := poly.Roots()
	if err != nil {
		t.Fatalf(`Roots() errored: %v`, err)
	}

	if len(roots) != 0 {
		t.Fatalf(`Roots() returned: %v for polynomial with no roots!`, roots)
	}

	fmt.Println("No Real Roots ......... OK")
}


func TestRealRoots1(t *testing.T){

	// a := 1.0
	// b := 1.5958982
	// c := -13.00789
	// d := -12.13037
	// e := 5.74231

	coeffs := []float64{-0.03998070290490681, 0.33278519947040125, -0.9247055429913947, 9.87899479042355, -0.09603364329060765}

	poly := CreatePolynomial(coeffs...)

	roots, err := poly.Roots()
	if err != nil {
		t.Fatalf(`Roots() errored: %v`, err)
	}

	sol1 := 0.009729823715772645
	sol2 := 8.858183272729985

	solutions := make(map[float64]bool)

	solutions[sol1] = true
	solutions[sol2] = true



	for _, root := range roots {

		for solution := range solutions {
			if root == Round(solution) {
				delete(solutions, solution)
				break
			}
		}
	}

    var resultStr string = "\n"
    for _, r := range roots {
        resultStr += fmt.Sprintf("%v\n", r)
    }

    solutionsSlice := []float64{sol1, sol2}
    var solutionStr string = "\n"
    for _, s := range solutionsSlice {
        solutionStr += fmt.Sprintf("%v\n", Round(s))
    }

	if len(solutions) != 0 {
		t.Fatalf(`Failed to find all correct real roots! 
		Found roots: %s. 
		Correct roots: %s`, 
		resultStr, 
		solutionStr)
	}

	fmt.Println("Real Roots ............ OK")
}



func TestDerivative(t *testing.T) {
	a := 1.0
	b := 1.0 
	c := 0.0
	d := -1.0 
	e := -1.0

	poly := CreatePolynomial(a, b, c, d, e)
	sol := []float64{4.0, 3.0, 0.0, -1.0}

	deriv := poly.Derivative()

	if deriv.Degree() != 3 || deriv.coeffs[0] != 4 || deriv.coeffs[1] != 3 || deriv.coeffs[2] != 0 || deriv.coeffs[3] != -1 {
		t.Fatalf(`Derivative() failed. Expected coeffs: %v. Received coeffs: %v`, sol, deriv.coeffs)
	}

	fmt.Println("Derivative ............ OK")
}

func TestComplexRoots(t *testing.T){
	
	a := 1.0
	b := 3.0
	c := -1.5
	d := -8.0
	e := -12.5

	poly := CreatePolynomial(a, b, c, d, e)

	roots, err := poly.ComplexRoots()
	if err != nil {
		t.Fatalf(`Roots() errored: %v`, err)
	}

	sol1 := complex(1.8892177902751495, 0)
	sol2 := complex(-3.071745756398733, 0)
	sol3 := complex(-0.9087360169382082, 1.152468686777906)
	sol4 := complex(-0.9087360169382082, -1.152468686777906)

	solutions := make(map[complex128]bool)
	solutions[sol1] = true
	solutions[sol2] = true
	solutions[sol3] = true
	solutions[sol4] = true

	for _, root := range roots {

		for solution := range solutions {
			if root == RoundC(solution) {
				delete(solutions, solution)
				break
			}
		}
	}

    var resultStr string = "\n"
    for _, r := range roots {
        resultStr += fmt.Sprintf("%v\n", r)
    }
    allSolutions := []complex128{sol1, sol2, sol3, sol4}
    var solutionStr string = "\n"
    for _, s := range allSolutions {
        solutionStr += fmt.Sprintf("%v\n", s)
    }

	if len(solutions) != 0 {
		t.Fatalf(`Failed to find all correct complex roots! 
		Found roots: %s. 
		Correct roots: %s`, 
		resultStr, 
		solutionStr)
	}

	fmt.Println("Complex Roots ......... OK")

}

func TestPoly(t *testing.T){
	a := 1.0
	b := 5.0
	c := 6.0

	evalPoint1 := 3.5
	evalPoint2 := complex(-2.3, 1.1)

	var solution1 float64 = 35.75
	solution2 := complex(-1.42, 0.44)

	test_poly := CreatePolynomial(a, b, c)

	result1 := test_poly.At(evalPoint1)

	if result1 != solution1 {
		t.Fatalf(`At() returned wrong solution: %f Expected: %f`, result1, solution1)
	}

	result2 := test_poly.AtComplex(evalPoint2)

	if result2 != solution2 {
		t.Fatalf(`At() returned wrong solution: %v Expected: %v`, result2, solution2)
	}

	fmt.Println("Polynomial ............ OK")

}

func TestQuadratic(t *testing.T) {
	a := 1.0
	b := 5.0
	c := 6.0

	solution1 := -2.0
	solution2 := -3.0


    test_poly := CreatePolynomial(a, b, c)

    roots, err := test_poly.Roots()
    if err != nil {
    	t.Fatalf(`Roots() errored: %v`, err)
    }

    var found1, found2 bool

    for _, root := range roots {
    	if root == solution1{
    		found1 = true
    	}
    	if root == solution2 {
    		found2 = true
    	}
    }

    if !found1 || !found2 {
    	t.Fatalf(`Roots() returned wrong solutions`)
    }


	fmt.Println("Quadratic Roots ....... OK")

}

