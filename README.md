![Logo](./images/golynomials.png)

# Polynomials for Go
### A Go package that handles the most essential polynomial operations

## Usage

Import by:   
```
import "github.com/mihonen/polynomials"

```



## Derivatives
A derivative can be obtained by: 
```
derivative := poly.Derivative()

```

## Root Solving
The package has three methods for solving roots of polynomials. 

| Method               | Complex Roots |  Average solve time<sup>1</sup>  | Robustness |
| -----------          | -----------   | --------------------------- |      -----------|
| Durand-Kerner        | ✅            | 6.623µs                     |         🥉     |
| Bisection + Newton   | ❌            | 7.38µs                      |         🥈     |
| Eigenvalue           | ✅            | 142.292µs                   |         🥇     |

<sup>1</sup> *Tested with 5 runs using polynomial:* $P(x) = 1.13x^4 - 5.0x^3 + 12.0x^2 -2.8x + 3.213$


The package uses Quadratic formula to solve roots for simple qudratic polynomials. The default method for higher order polynomials computes the companion matrix of the polynomial and finds the eigenvalues of the matrix using [mat package](https://pkg.go.dev/gonum.org/v1/gonum/mat). 



The second available method is a combination of bisection method and Newton-method as described in [this](https://en.wikipedia.org/wiki/Real-root_isolation#Bisection_method) and this [page](https://en.wikipedia.org/wiki/Sturm%27s_theorem#Root_isolation). This method first utilizes [Sturm's theorem](https://en.wikipedia.org/wiki/Sturm%27s_theorem) to seek for intervals which hold exactly one real root. It then finds the roots numerically using [Newton's-method](https://en.wikipedia.org/wiki/Newton%27s_method).  
    
    

The third available method is the [Durand-Kerner method](https://en.wikipedia.org/wiki/Durand–Kerner_method). This method should be able to solve all complex roots for polynomials upto around 100 degrees. 
    
    

Used method can be changed by changing the field SolveMode of the polynomial. For example
```
poly.SolveMode = polynomials.DurandKerner

```


### Getting Complex Roots

```
roots, err := poly.ComplexRoots()

```

### Getting Real Roots

```
roots, err := poly.Roots()

```


## Examples 
### Solving Complex Roots for $P(x) = 3x^3 + 2x^2 -x + 13$


```
    a :=  3.0
    b :=  2.0
    c := -1.0
    d :=  13.0

    poly := polynomials.CreatePolynomial(a, b, c, d)

    roots, err := poly.ComplexRoots()
    if err != nil {
        log.Printf(`ComplexRoots() errored: %v`, err)
        return
    }
    // Use roots

```

### Solving Real Roots for $P(x) = x^4 + 2x^2 -10$


```
    a :=  1.0
    b :=  0.0
    c :=  2.0
    d :=  0.0
    e := -10.0

    poly := polynomials.CreatePolynomial(a, b, c, d, e)

    roots, err := poly.Roots()
    if err != nil {
        log.Printf(`Roots() errored: %v`, err)
        return
    }
    // Use roots

```


## Precision

The package solves roots to the 9th decimal by default. This can be adjusted in config.go if needed.


## Testing

To run all tests, run the following command in terminal
```
go test

```

## TODO


- [ ] Check that root returned by Newton's method lies in the given interval

- [ ] Add functionalities for solving minimums and maximums

- [x] Implement more robust complex root finding algorithm, that finds the eigenvalues of the companion matrix

## Credits
This project is partly based on [polygo](https://github.com/SeanJxie/polygo) by [Sean Xie](https://github.com/SeanJxie). We acknowledge and appreciate their work in the initial stages of this project.


