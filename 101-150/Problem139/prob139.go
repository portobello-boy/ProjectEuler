package main

import (
	"fmt"
	"math/big"
)

func getTriangles() int {
	count := 0
	for n := 1; n < 4083; n++ {
		for m := n + 1; m < 5774; m += 2 {
			// Ensure m, n are coprime
			var g big.Int
			g.GCD(nil, nil, big.NewInt(int64(m)), big.NewInt(int64(n)))
			if g.Cmp(big.NewInt(1)) != 0 {
				continue
			}

			// Generate a, b, c
			a := (m * m) - (n * n)
			b := 2 * m * n
			c := (m * m) + (n * n)

			if c%(b-a) != 0 {
				continue
			}

			sum := a + b + c
			k := 1

			for k*sum < 100000000 {
				ka, kb, kc := k*a, k*b, k*c

				diff := kb - ka

				if ka == 5 && kb == 12 && kc == 13 && kc%diff == 0 {
					print(ka, kb, kc, diff)
				}

				if kc%diff == 0 {
					count++
				}

				k++
			}
		}
	}
	return count
}

func main() {
	fmt.Println(getTriangles())
}
