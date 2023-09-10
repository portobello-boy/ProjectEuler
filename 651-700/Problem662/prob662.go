package main

import (
	"MathLibGo/modulo"
	"fmt"
	"numberTheory/sequences"
	"numberTheory/tools"

	"golang.org/x/exp/slog"
)

var fibonacciSequence []int
var fibonacciSequenceLen int
var exploredGrids map[string]int = make(map[string]int)

func F(W, H int, adder func(int) int) int {
	if paths, ok := exploredGrids[fmt.Sprintf("%d-%d", W, H)]; ok {
		adder(paths)
		return paths
	}
	if W == 0 && H == 0 {
		adder(1)
		return 1
	} else if W < 0 || H < 0 {
		return 0
	}

	totalPaths, _ := modulo.ModularAdder(1000000007)

	// totalPaths := 0
	maxDistance := tools.HypotenuseInt(W, H)

	totalPaths(F(W-1, H, adder))
	totalPaths(F(W, H-1, adder))

	for _, f := range fibonacciSequence {
		if f == 0 || f == 1 { // 1 is accounted for above
			continue
		}
		if f > maxDistance {
			break
		}
		if f <= W {
			totalPaths(F(W-f, H, adder))
		}
		if f <= H {
			totalPaths(F(W, H-f, adder))
		}
	}

	// To determine which Fibonacci numbers are Pythagorean triples, refer to this information
	// https://en.wikipedia.org/wiki/Pythagorean_triple#Fibonacci_numbers_in_Pythagorean_triples
	n := 1

	for 2*n+3 <= fibonacciSequenceLen {
		if fibonacciSequence[2*n+3] > maxDistance {
			break
		}
		p := fibonacciSequence[n] * fibonacciSequence[n+3]
		q := 2 * fibonacciSequence[n+1] * fibonacciSequence[n+2]
		totalPaths(F(W-p, H-q, adder))
		totalPaths(F(W-q, H-p, adder))
		n++
	}

	exploredGrids[fmt.Sprintf("%d-%d", W, H)] = totalPaths(0)
	return totalPaths(0)
}

func solveProblem(W, H int) int {
	adder, _ := modulo.ModularAdder(1000000007)

	fibs, _ := sequences.BoundedFibonacciSequence(tools.HypotenuseInt(W, H) + 1)
	fibonacciSequence = fibs
	fibonacciSequenceLen = len(fibonacciSequence)

	slog.Info("Ready for processing")
	fmt.Println(F(W, H, adder))
	return adder(0)

}

func main() {
	fmt.Println(solveProblem(10000, 10000))
}
