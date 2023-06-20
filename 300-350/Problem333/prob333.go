package main

import (
	"fmt"
	"math"
	"sort"
	"sync"
	"time"
)

const THREADS = 200
const BOUND = 1000000

func isPrime(num int) bool {
	if num%2 == 0 || num%3 == 0 {
		return false
	}

	for k := 1; (6*k-1)*(6*k-1) <= num; k++ {
		if num%(6*k-1) == 0 || num%(6*k+1) == 0 {
			return false
		}
	}

	return true
}

func generatePrimesInBound(bound int) []int {
	// So we don't allocate needless amounts of memory, we make an array with a known upper
	// bound on pi(x) for x >= 17
	primes := make([]int, int(math.Round(1.25506*(float64(bound)/math.Log(float64(bound))))))
	primes[0] = 2
	primes[1] = 3

	index := 2         // How many primes have we encountered
	k := 1             // Multiple of 6
	possiblePrime := 3 // What number are we looking at

	for possiblePrime < bound {
		possiblePrime = 6*k - 1
		if possiblePrime >= bound { // Make sure we aren't past the bound
			break
		}

		if isPrime(possiblePrime) {
			primes[index] = possiblePrime
			index++
		}

		possiblePrime += 2
		if possiblePrime >= bound { // This statement is needed in the event that the desired prime is a twin prime
			break
		}

		if isPrime(possiblePrime) {
			primes[index] = possiblePrime
			index++
		}
		k++
	}

	return primes[:index]
}

func decomposeNumber(num int, currentPartitions, partitionTerms *[]int) int {
	// fmt.Println(num, currentPartitions, partitionTerms)

	if num == 0 {
		return 1
	}

	if num < 0 || num == 1 {
		return 0
	}

	valid := 0
	for i := len(*partitionTerms) - 1; i >= 0; i-- {
		termToCheck := (*partitionTerms)[i]

		// fmt.Printf("Term to check: %d, current partition: %v\n", termToCheck, currentPartitions)

		// If the term to check divides or is divisble by anything in currentPartition, skip this term
		validTerm := true
		for _, term := range *currentPartitions {

			// term should always be larger than termToCheck, by algorithm design
			if term%termToCheck == 0 {
				// fmt.Printf("Found divisible terms: %d mod %d == 0\n", termToCheck, term)
				validTerm = false
				break
			}
		}

		if !validTerm {
			continue
		}

		newCurrentPartition := append(*currentPartitions, termToCheck)
		usableTerms := (*partitionTerms)[:i]

		// Recurse through different terms and find all valid partitions
		valid += decomposeNumber(num-termToCheck, &newCurrentPartition, &usableTerms)

		// As soon as we find a second valid partition, just return
		if valid >= 2 {
			// fmt.Printf("Found multiple valid partitions, returning...\n")
			return valid
		}
	}
	return valid
}

// A structure to hold the necessary data for a goroutine worker
type Job struct {
	Id        int
	Prime     int
	TermIndex int
}

func solveProblem(bound int) int {
	// Generate the necessary primes and valid terms under the provided bound
	primes := generatePrimesInBound(bound)
	terms := generateTerms(bound)

	// primesLen represents the total number of jobs needed
	primesLen := len(primes)
	fmt.Printf("Generated %d primes\n", primesLen)
	termsLen := len(terms)
	fmt.Printf("Generated %d terms\n", termsLen)

	// Initiate the total sum counter
	sum := 0

	// Create helpers for tracking time and worker pool
	startTime := time.Now()
	var wg sync.WaitGroup

	jobs := make(chan Job, primesLen)
	results := make(chan int, THREADS)
	done := make(chan bool)

	// Create a function to handle listening to the jobs channel and partitioning primes
	decompose := func(threadId int, terms []int) {
		defer wg.Done()
		for job := range jobs {

			// Capture the time informatin for every 10th job
			var jobStartTime, jobEndTime time.Time
			if job.Id%10 == 0 {
				jobStartTime = time.Now()
			}

			// For some reason, slices can't be passed by reference within function calls, so it needs to be allocated here before passing as a parameter
			usableTerms := terms[:job.TermIndex+1]

			// Begin recursive call to find valid partitions of a prime
			decompositions := decomposeNumber(job.Prime, &[]int{}, &usableTerms)

			if job.Id%10 == 0 {
				jobEndTime = time.Now()
				fmt.Printf("ThreadId: %d Job: %+v Total Time: %f seconds\n", threadId, job, jobEndTime.Sub(jobStartTime).Seconds())
			}

			// Pass the decomposition data back to the results channel
			if decompositions == 1 {
				results <- job.Prime
			} else {
				results <- 0
			}
		}
	}

	// Create a function to read all results and calculate the answer
	getResults := func() {
		for val := range results {
			sum += val
		}
		done <- true
	}

	// Allocate the workload to the jobs channel and close it
	termIndex := 0
	for i, prime := range primes {
		for termIndex+1 < termsLen && terms[termIndex+1] <= prime {
			termIndex += 1
		}

		job := Job{
			i, prime, termIndex,
		}

		jobs <- job
	}
	close(jobs)

	// Allocate goroutine to get results
	go getResults()

	// Start goroutines to process, wait for them to complete
	for thread := 0; thread < THREADS; thread++ {
		wg.Add(1)
		go decompose(thread, terms)
	}

	// Wait for the calculating threads to finish and then close the results channel
	wg.Wait()
	close(results)

	// Wait for the completion signal
	<-done

	endTime := time.Now()
	fmt.Printf("Total time: %f seconds\n", endTime.Sub(startTime).Seconds())

	return sum
}

func generateTerms(bound int) []int {
	var terms []int

	x := 0
	for i := 1; i < bound; i = int(math.Pow(2, float64(x))) {
		y := 0
		for j := 1; j < bound; j = int(math.Pow(3, float64(y))) {
			product := i * j
			if product < BOUND {
				terms = append(terms, product)
			} else {
				break
			}
			y++
		}
		x++
	}

	sort.Ints(terms)
	return terms[1:]
}

func main() {
	fmt.Println(solveProblem(BOUND))
}
