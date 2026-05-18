package main

import (
	helper "MathLibGo"
	"MathLibGo/operations/combinatorics"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var digits = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(filename string) map[int][]string {

	f, err := os.Open(filename)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()

	str := scanner.Text()
	str = strings.TrimSuffix(str, "\n")
	rowStr := strings.Split(str, ",")

	lenToWordList := make(map[int][]string)

	for _, val := range rowStr {
		val = strings.Trim(val, "\"")
		lenToWordList[len(val)] = append(lenToWordList[len(val)], val)
	}

	return lenToWordList
}

var cached_IsSquare = helper.CacheWrapper_SingleToSingle(func(n int) bool {
	sqrt := math.Sqrt(float64(n))
	return sqrt == float64(int(sqrt))
})

func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	aMap := make(map[rune]int)
	for _, r := range a {
		aMap[r]++
	}
	for _, r := range b {
		aMap[r]--
		if aMap[r] < 0 {
			return false
		}
	}
	return true
}

func arePalindromic(a, b string) bool {
	for i := 0; i < len(a)/2; i++ {
		if a[i] != b[len(b)-1-i] {
			return false
		}
	}
	return true
}

func checkSolution(word, otherWord string, d []rune) int {
	maxNum := 0

	// We do not want leading zeroes in the number
	if d[0] == '0' {
		return 0
	}
	charToDigit := make(map[rune]rune)
	concatenatedNumber := ""

	for i, r := range word {
		charToDigit[r] = d[i]
		concatenatedNumber += string(d[i])
	}

	num, _ := strconv.Atoi(concatenatedNumber)
	if !cached_IsSquare(num) {
		return 0
	}

	maxNum = max(maxNum, num)

	concatenatedNumber = ""
	for _, r := range otherWord {
		concatenatedNumber += string(charToDigit[r])
	}

	// We do not want leading zeroes in the number
	if concatenatedNumber[0] == '0' {
		return 0
	}

	num, _ = strconv.Atoi(concatenatedNumber)
	if !cached_IsSquare(num) {
		return 0
	}

	return max(maxNum, num)
}

func main() {
	words := readFile("./words.txt")
	fmt.Println(words)

	for wordLen, wordList := range words {
		for j, word := range wordList {
			// Find anagrams
			for k, otherWord := range wordList {
				if j != k && isAnagram(word, otherWord) && !arePalindromic(word, otherWord) {
					fmt.Printf("%s and %s are anagrams\n", word, otherWord)
					// Assign a digit to each letter and determine if they are squares
					// Permute digits 0-9 of length wordLength and check if they are squares
					digitPermutations := combinatorics.Permute(digits, wordLen)

					for _, dp := range digitPermutations {
						solution := checkSolution(word, otherWord, dp)
						if solution != 0 {
							fmt.Println(solution)
						}
					}

				}
			}
		}
	}

	fmt.Println(cached_IsSquare(1296))
	fmt.Println(cached_IsSquare(1296))

	fmt.Println(isAnagram("CARE", "RACE"))

	fmt.Println(combinatorics.Permute([]int{1, 2, 3, 4}, 3))

}
