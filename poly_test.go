package polynomials


import (
    "testing"

)



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

}

func TestQudratic(t *testing.T) {
	a := 1.0
	b := 5.0
	c := 6.0

	solution1 := -2.0
	solution2 := -3.0


    test_poly := CreatePolynomial(a, b, c)

    roots, err := test_poly.PositiveRoots()
    if err != nil {
    	t.Fatalf(`PositiveRoots() errored: %v`, err)
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

