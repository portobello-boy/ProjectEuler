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
	"regexp"

	"gonum.org/v1/gonum/stat/combin"
)

// Observations:
// (2k choose k) is always even
// S(n) can be generated more easily using previous terms
// 3S(n) + 4 guarantees that v2(result) is > 2

var re *regexp.Regexp = regexp.MustCompile(`^[01\s]*(110|001)[01\s]*100$`)

func numToBinaryString(num int64, addSpace bool) string {
	bin := fmt.Sprintf("%032b", uint32(num))

	if addSpace {
		for i := 4; i < len(bin); i += 5 {
			bin = bin[:i] + " " + bin[i:]
		}
	}
	return bin
}

func altv2(n int64) int64 {
	snBinaryNoPad := numToBinaryString(n, false)

	regexmatch := re.Match([]byte(snBinaryNoPad))

	ret := int64(0)

	if regexmatch {
		matches := re.FindStringSubmatchIndex(snBinaryNoPad)
		ret = int64((len(snBinaryNoPad) - matches[2]) - 1)
	}

	return ret
}

func v2(n int64) int64 {
	return int64(bits.TrailingZeros64(uint64(n)))
}

func S(n int64) int64 {
	sum := int64(0)
	for k := int64(1); k <= n; k++ {
		pow2 := int64(math.Pow(-2, float64(k)))
		binom := int64(combin.Binomial(int(2*k), int(k)))
		term := (pow2 * binom)

		sum += term

		// fmt.Printf("S(%d) - k: %2d, term: %10d, binary: %s, binomial term: %6d, binomial binary: %s\n", n, k, term, numToBinaryString(term), binom, numToBinaryString(binom))
		fmt.Printf("S(%d) - k: %2d, binomial term: %6d, binomial binary: %s\n", n, k, binom, numToBinaryString(binom, true))
		fmt.Printf("                       Term: %10d, binary term: %s\n", term, numToBinaryString(term, true))
		fmt.Printf("                  Sum so far: %10d, binary sum: %s\n", sum, numToBinaryString(sum, true))
		fmt.Printf("\n")
	}

	fmt.Printf("                       S(%d): %10d, binary S(%d): %s\n", n, sum, n, numToBinaryString(sum, true))
	fmt.Printf("                 3*(S(%d)+4): %10d, binary     : %s\n", n, (3*sum)+4, numToBinaryString((3*sum)+4, true))
	fmt.Printf("                                               v2(x): %10d\n", v2((3*sum)+4))
	fmt.Printf("                                  alt v2(Sum so far): %10d\n", altv2(sum))
	return sum
}

func u(n int64) int64 {
	sn := S(n)
	snBinary := numToBinaryString(sn, true)
	snBinaryNoPad := numToBinaryString(sn, false)
	fmt.Printf("S(%d): %10d, binary S(%d): %s\n", n, sn, n, snBinary)
	fmt.Printf("                           <<: %s\n", numToBinaryString((sn<<1), true))
	fmt.Printf("                       3*S(%d): %s\n", n, numToBinaryString(3*sn, true))
	fmt.Printf("                     3*S(%d)+4: %s\n", n, numToBinaryString((3*sn)+4, true))
	fmt.Printf("                 v2(3*S(%d)+4): %d\n", n, v2((3*sn)+4))
	// fmt.Printf("S(%d): %10d, binary S(%d): %s\n", n, sn, n, snBinary)

	pattern := `^[01\s]*(110|001)[01\s]*100$`
	re := regexp.MustCompile(pattern)

	regexmatch := re.Match([]byte(snBinaryNoPad))
	fmt.Printf("%s matches pattern %s: %t\n", snBinary, pattern, regexmatch)

	if regexmatch {
		matches := re.FindStringSubmatchIndex(snBinaryNoPad)
		fmt.Println((len(snBinaryNoPad) - matches[2]) - 1)
	}

	return v2((3 * sn) + 4)
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

	for i := 1; i <= 5; i++ {
		fmt.Println(S(int64(i)))
	}

	// fmt.Println(S(8))
	// fmt.Println(v2(24))
	// fmt.Println(S(4))
	// fmt.Println(u(6))
	// fmt.Println(U(2))
}
