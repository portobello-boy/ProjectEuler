package main

import (
	"MathLibGo/modulo"
	"fmt"
	"math"
)

var NUM_DIGITS float64 = 10

func main() {
	fmt.Println("Hello World")
	multr, _ := modulo.ModularMultiplier(int(math.Pow(10, NUM_DIGITS)))
	res := multr(28433)
	for i := 0; i < 7830457; i++ {
		res = multr(2)
	}
	res += 1
	fmt.Println(res)
}
