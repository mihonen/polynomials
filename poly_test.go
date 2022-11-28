package polynomials


import (
	"fmt"
    "testing"
)




func TestRealRoots(t *testing.T){

	a := 1.0
	b := 1.5958982
	c := -13.00789
	d := -12.13037
	e := 5.74231


	poly := CreatePolynomial(a, b, c, d, e)

	roots, err := poly.Roots()
	if err != nil {
		t.Fatalf(`Roots() errored: %v`, err)
	}

	fmt.Println("Roots() OK")
}


func TestCreatePower(t *testing.T){
	x_3 := CreatePower(3)
	if len(x_3.coeffs) != 4 || x_3.coeffs[0] != 1.0 {
		t.Fatalf(`CreatePower() failed`,)
	}

	fmt.Println("CreatePower() OK")
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

	fmt.Println("Derivative() OK")
}

func TestDurandKerner(t *testing.T){
	
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
		t.Fatalf(`Roots() failed to find all correct roots! 
		Found roots: %s. 
		Correct roots: %s`, 
		resultStr, 
		solutionStr)
	}

	fmt.Println("DurandKerner() OK")

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

	fmt.Println("Polynomial OK")

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
}

