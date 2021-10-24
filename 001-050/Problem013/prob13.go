package main

import (
	"os"
	"fmt"
	"strings"
	"io/ioutil"
	"github.com/holiman/uint256"
)

func buildLargeNumArray(text string) []*uint256.Int {
	numsStr := strings.Split(text, "\n")
	numsInt := make([]*uint256.Int, len(numsStr))

	for i, num := range numsStr {
		v, err := uint256.FromHex(num)
		if err != nil {
			fmt.Println(err)
		}

		numsInt[i] = v
	}

	return numsInt
}

func addNums(nums []*uint256.Int) *uint256.Int {
	sum := uint256.NewInt()

	for _, num := range nums {
		sum.Add(sum, num)
	}

	return sum
}

// Thanks to https://www.convzone.com/decimal-to-hex/ for converting enormous numbers
// Thanks to https://github.com/holiman/uint256 for dealing with stupidly huge numbers
func main() {
	file := os.Args[1]

	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}

	nums := buildLargeNumArray(string(content))
	fmt.Println(addNums(nums))
}