package main

import (
	mlg "MathLibGo"
	"fmt"
	"time"
)

func f1(v int) int {
	time.Sleep(500 * time.Millisecond)
	return v
}

func f2(v1 int, v2 ...int) int {
	time.Sleep(500 * time.Millisecond)
	// v2 = append(v2, v1)
	// s, _ := json.Marshal(v2)
	return v1
	// return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(v2)), ", "), "[]")
}

func f3(v1 int, v2 string, v3 ...int) int {
	time.Sleep(500 * time.Millisecond)

	fmt.Println(v2)
	// v2 = append(v2, v1)
	// s, _ := json.Marshal(v2)
	return v1
	// return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(v2)), ", "), "[]")
}

func main() {
	// cachedFunc := mlg.CacheWrapper_SingleToSingle(f1)
	cachedFunc := mlg.CacheWrapper_ManyToSingle(f2)

	fmt.Println(cachedFunc(1, 2, 3))
	fmt.Println(cachedFunc(1, 2, 3))
	fmt.Println(cachedFunc(2, 3, 2))
	fmt.Println(cachedFunc(2, 3))
	fmt.Println(cachedFunc(3, 3))
	fmt.Println(cachedFunc(2, 3))

	// for i := 0; i < 100; i++ {
	// 	fmt.Println(cachedFunc(rand.Intn(100 - 0 + 1)))
	// }
}
