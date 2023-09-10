package figurate

import "errors"

// Returns nth triangular number
func Triangular(n int) (int, error) {
	if n < 0 {
		return -1, errors.New("n must be a Natural number")
	}
	return n * (n + 1) / 2, nil
}

// Returns nth polygonal number of dimension r
// Triangular : r=3
func Polygonal(r, n int) (int, error) {
	if n < 0 {
		return -1, errors.New("n must be a Natural number")
	}
	if r < 2 {
		return -1, errors.New("r must be a Natural number greater than 1")
	} else if r == 2 {
		return n, nil
	}

	nMinusOneTriangular, _ := Triangular(n - 1)
	nTriangular, _ := Triangular(n)
	return (r-3)*nMinusOneTriangular + nTriangular, nil
}
