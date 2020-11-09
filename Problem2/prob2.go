package main

import (
	"os"
	"fmt"
	"strconv"
	"math"
)

// Using closures, this problem can be solved by calculating all Fibonacci terms
// up to 4,000,000. It's a brute force method, but using closures in Golang for
// what is essentially 'functional memory' is pretty awesome.
func approach1(num int) {
	sum := 0
	nextFib := fibSeq()
	term := nextFib()

	for term < num {
		if term % 2 == 0 {
			sum += term
		}

		term = nextFib()
	}

	fmt.Println(sum)
}

func fibSeq() func() int {
	i := 0
	j := 1
	return func() int {
		k := i + j
		i = j
		j = k
		return k
	}
}

// Approach 2 uses some analysis. The ratio of two sequential Fibonacci numbers approaches
// Phi, the golden ratio. Even-values Fibonacci terms are separated by two odd terms, so
// each even term is approximately 3 golden ratios apart. We can use this to calculate the sum of terms
// without generating each Fibonacci value.
func approach2(num int) {
	sum := 0
	term := 2

	for term < num {
		sum += term
		term = int(math.Round(float64(term) * math.Pow(math.Phi, 3)))
	}

	fmt.Println(sum)
}

func main() {
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	// approach1(num)
	approach2(num)
}