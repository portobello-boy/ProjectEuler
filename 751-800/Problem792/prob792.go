// We define $\nu_2(n)$ to be the largest integer $r$ such that $2^r$ divides $n$. For example, $\nu_2(24) = 3$.

// Define $\displaystyle S(n)  = \sum_{k = 1}^n (-2)^k\binom{2k}k$ and $u(n) = \nu_2\Big(3S(n)+4\Big)$.

// For example, when $n = 4$ then $S(4) = 980$ and $3S(4) + 4 = 2944 = 2^7 \cdot 23$, hence $u(4) = 7$.
// You are also given $u(20) = 24$.

// Also define $\displaystyle U(N) = \sum_{n = 1}^N u(n^3)$. You are given $U(5) = 241$.

// Find $U(10^4)$.

package main

import (
	"fmt"
	"math"
	"math/bits"

	"gonum.org/v1/gonum/stat/combin"
)

func v2(n int64) int64 {
	return int64(bits.TrailingZeros64(uint64(n)))
}

func S(n int64) int64 {
	sum := int64(0)
	for k := int64(1); k <= n; k++ {
		sum += (int64(math.Pow(-2, float64(k))) * int64(combin.Binomial(int(2*k), int(k))))
	}
	return sum
}

func u(n int64) int64 {
	sn := S(n)
	fmt.Println(sn)
	return v2((3 * S(n)) + 4)
}

func U(N int) int64 {
	sum := int64(0)
	for n := 1; n <= N; n++ {
		fmt.Println(n, math.Pow(float64(n), 3))
		sum += u(int64(math.Pow(float64(n), 3)))
	}
	return sum
}

func main() {
	fmt.Println(S(8))
	// fmt.Println(v2(24))
	// fmt.Println(S(4))
	// fmt.Println(u(4))
	// fmt.Println(U(5))
}
