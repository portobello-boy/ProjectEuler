package main

import (
	helper "MathLibGo"
	"MathLibGo/numberTheory/primes"
	"MathLibGo/operations/factorial"
	"fmt"
	"iter"
	"math"
	"sort"
	"sync"
)

var BOUND uint64 = 31
var CALCULATION_BOUND uint64 = factorial.DoubleFactorial(BOUND)

var hashToSequence = map[uint64]*Sequence{}

// var hashMapToValue = map[uint64][]uint64{}
var valueToHash = map[uint64]uint64{}
var hashMapIndex = uint64(2)
var mapLock = sync.Mutex{}

var optimalFactor = helper.CacheWrapper_SingleToSingle(func(value uint64) uint64 {
	if primes.IsProbablePrime(value) {
		return 1
	}
	sqrt := uint64(math.Sqrt(float64(value)))
	for n := sqrt; n > 1; n-- {
		if value%n == 0 {
			return n
		}
	}
	return 1
})

type Tree struct {
	val   uint64
	left  *Tree
	right *Tree
}

func newLeaf(value uint64) *Tree {
	return &Tree{
		value,
		nil,
		nil,
	}
}

func buildTree(rootValue uint64) *Tree {
	root := newLeaf(rootValue)
	left := optimalFactor(rootValue)
	right := rootValue / left

	if !(left == 1 && right == rootValue) {
		root.left = buildTree(left)
		root.right = buildTree(right)
	}

	return root
}

func (t *Tree) copy() *Tree {
	if t.isLeaf() {
		return newLeaf(t.val)
	}
	return &Tree{
		t.val,
		t.left.copy(),
		t.right.copy(),
	}
}

func (t *Tree) isLeaf() bool {
	return t.left == nil && t.right == nil
}

func (t *Tree) countLeaves() uint64 {
	if t.isLeaf() {
		return 1
	}
	return t.left.countLeaves() + t.right.countLeaves()
}

func (t *Tree) Print() {
	fmt.Printf("%v\n", t)
	if t.left != nil {
		t.left.Print()
	}
	if t.right != nil {
		t.right.Print()
	}
}

func (t *Tree) hash() uint64 {
	mapLock.Lock()
	defer mapLock.Unlock()

	frontier := []*Tree{t}
	hash := uint64(0)

	if h, ok := valueToHash[t.val]; ok {
		return h
	}

	for len(frontier) > 0 {
		next := []*Tree{}
		hasNext := false
		for _, node := range frontier {
			hash = hash << 1
			if node != nil {
				hash += 1
				next = append(next, node.left, node.right)
				hasNext = true
				// } else {
				// 	next = append(next, nil, nil)
			}
		}
		frontier = next
		if !hasNext {
			break
		}
	}

	for hash%2 == 0 {
		hash = hash >> 1
	}

	valueToHash[t.val] = hash
	// hashMapToValue[hash] = append(hashMapToValue[hash], t.val)

	return hash
}

type Sequence struct {
	hash      uint64           // Hash of tree structure defining this sequence
	values    map[uint64]bool  // Ordered list of values already explored for this sequence
	generator iter.Seq[uint64] // Sequence generator for this sequence
}

func (s *Sequence) iterateSequence(lowerBound uint64) iter.Seq[uint64] {
	return func(yield func(i uint64) bool) {
		// Sort the keys of the values
		keys := make([]uint64, 0)
		for k, _ := range s.values {
			// if k < lowerBound {
			// 	continue
			// }
			keys = append(keys, k)
		}
		sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

		for _, k := range keys {
			// for k := range s.values {
			// if k > max {
			// 	return
			// }
			if !yield(k) {
				return
			}
		}

		for newValue := range s.generator {

			// s.values = append(s.values, newValue)
			if _, ok := s.values[newValue]; !ok {
				s.values[newValue] = true
			}
			// fmt.Println("NEW VALUE")
			// if newValue > max {
			// 	return
			// }
			// if newValue < lowerBound {
			// 	continue
			// }

			if !yield(newValue) {
				return
			}
		}
	}
}

func getTreeShape(n uint64) uint64 {

	mapLock.Lock()
	defer mapLock.Unlock()

	frontier := []uint64{n}
	hash := uint64(0)

	if h, ok := valueToHash[n]; ok {
		return h
	}

	for len(frontier) > 0 {
		next := []uint64{}
		hasNext := false
		for _, f := range frontier {
			hash <<= 1
			if f == 0 {
				continue
			}
			leftFactor := optimalFactor(f)
			rightFactor := f / leftFactor

			hash += 1

			if leftFactor != 1 {
				// getTreeShape(leftFactor)
				// getTreeShape(rightFactor)
				next = append(next, leftFactor, rightFactor)
				hasNext = true
			} else {
				next = append(next, 0, 0)
			}
		}
		frontier = next
		if !hasNext {
			break
		}
	}

	for hash%2 == 0 {
		hash = hash >> 1
	}

	valueToHash[n] = hash
	// hashMapToValue[hash] = append(hashMapToValue[hash], n)

	return hash
}

func generatePrimes() iter.Seq[uint64] {
	return func(yield func(i uint64) bool) {
		p := uint(1)
		for {
			p, _ = primes.NextPrime(p)

			if p > uint(CALCULATION_BOUND) {
				return
			}

			if !yield(uint64(p)) {
				return
			}
		}
	}
}

// func (t *Tree) solver(ctx context.Context) <-chan uint64 {
func (t *Tree) solver() iter.Seq[uint64] {
	// sequence := make(chan uint64)
	// defer close(sequence)

	if seq, ok := hashToSequence[t.hash()]; ok {
		return seq.iterateSequence()
	}

	// Create goroutine to produce to channel
	if t.isLeaf() {
		// Push primes to channel
		if seq, ok := hashToSequence[1]; ok {
			return seq.iterateSequence()
		}
		primes := generatePrimes()
		return primes
	}
	// If we've already found a good shape for the current tree, return it
	// Else, generate it
	return func(yield func(i uint64) bool) {

		if t.val == 60 {
			fmt.Println("HERE")

		}
		rightVals := t.right.solver()

		// TODO - Priority Queue

		for r := range rightVals {
			if r > CALCULATION_BOUND {
				return
			}

			leftVals := t.left.solver()

		inner:
			for l := range leftVals {
				if l > r {
					break inner
					// continue
				}

				product := l * r

				if product > CALCULATION_BOUND {
					break inner
					// continue
				}

				if l != optimalFactor(product) {
					continue
				}

				if getTreeShape(product) != t.hash() {
					continue
				}

				if !yield(product) {
					return
				}
			}
		}
	}
}

func M(n uint64) []uint64 {

	solutions := []uint64{}

	t := buildTree(factorial.DoubleFactorial(n))
	for v := range t.solver() {
		solutions = append(solutions, v)
	}

	return solutions
}

func getGenerator(t *Tree, m map[uint64]*Sequence) {
	th := t.hash()
	if _, ok := m[th]; !ok {
		m[th] = &Sequence{
			hash:      th,
			values:    map[uint64]bool{},
			generator: t.solver(),
		}
	}

	if !t.isLeaf() {
		getGenerator(t.left, m)
		getGenerator(t.right, m)
	}
}

func main() {
	/*
		 1. Iterate through all targets
		 	- generate the tree
		 	- record the hash of the tree as a target
		  - discard duplicates since we only need the min value per shape
			- result: []*Tree{...list of unique target tree structures...},
					  map[uint64]uint64{ map n -> hash of n!! }
		 2. For each tree
		 	- map hash of tree to a generator function for that hash
			- recurse through children nodes for each tree and repeat
			- result: map[uint64]func() uint64{ map hash -> generator function for values }
		 			  map[uint64][]uint64{ map hash -> sequence of values (ordered) }
		 3, For each target hash
			- Collect the sequence for the target hash
			- Take the sum of the 0th index for the sequence of each hash

		Node Value generators:
			- Leaves emit sequence of primes to channel
			- Non-leaves take product of values from child nodes, and:
				- validate result is ideal
					- l < r, product has optimal shape
				- push result to internal priority queue
				- ensure results are ordered
				- emit queue's next item
	*/

	targetTrees := make([]*Tree, 0)
	targetToTreeHash := make(map[uint64]uint64)

	for n := uint64(2); n <= BOUND; n++ {
		t := buildTree(factorial.DoubleFactorial(n))
		tHash := t.hash()
		targetTrees = append(targetTrees, t)
		targetToTreeHash[n] = tHash
	}

	for _, t := range targetTrees {
		getGenerator(t, hashToSequence)
	}

	for k, v := range targetToTreeHash {
		fmt.Printf("%d!! = %d - %b - %v\n", k, factorial.DoubleFactorial(k), v, hashToSequence[v])
	}

	// s := Sequence{
	// 	hash:      0,
	// 	values:    []uint64{2, 3},
	// 	generator: generatePrimes(),
	// }

	// fmt.Println(s.values)

	// for p := range s.iterateSequence() {
	// 	fmt.Println(p)

	// 	if p > 20 {
	// 		break
	// 	}
	// }
	// fmt.Println(getTreeShape(54))

	// vals := make([]uint64, 0)
	// for s := range hashToSequence[targetToTreeHash[10]].iterateSequence(targetTrees[10-2].val) {
	// 	vals = append(vals, s)
	// 	// break
	// }
	// fmt.Printf("%d!! = %d - %d\n", 10, factorial.DoubleFactorial(10), vals)

	// fmt.Println(hashToSequence[127])

	// fmt.Println(s.values)
	// fmt.Println(hashToSequence)
	// // Sort the keys of the values
	// keys := make([]uint64, 0)
	// for k, _ := range hashToSequence {
	// 	keys = append(keys, k)
	// }
	// sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	// for _, k := range keys {
	// 	fmt.Println("generating", k)
	// 	counter := 0
	// 	for _ = range hashToSequence[k].iterateSequence() {
	// 		counter += 1
	// 		if counter%10000 == 0 {
	// 			fmt.Println(len(hashToSequence[k].values))

	// 		}
	// 	}
	// }
	// fmt.Println(hashToSequence)

	for n := uint64(2); n <= BOUND; n++ {
		// tm := time.Now()
		vals := make([]uint64, 0)
		// val := uint64(0)
	inner:
		for s := range hashToSequence[targetToTreeHash[n]].iterateSequence() {
			vals = append(vals, s)
			if n < 26 {
				if s > factorial.DoubleFactorial(n-1) {
					break inner
				}
			} else {
				break inner
			}
			// if len(vals) > 2 {
			// 	break
			// }
			// val = s
			// break
		}
		fmt.Printf("%d!! = %d - %d\n", n, factorial.DoubleFactorial(n), vals)
		// fmt.Println(hashToSequence[1].values)
	}

	// for k, s := range hashToSequence {
	// 	fmt.Printf("[%b - ", k)
	// 	for v := range maps.Keys(s.values) {
	// 		fmt.Printf("%d, ", v)
	// 	}
	// 	fmt.Println("]")

	// 	// fmt.Println(k, s.values)

	// }

	// fmt.Println(hashToSequence[targetToTreeHash[7]])

	// for s := range hashToSequence[targetToTreeHash[7]].iterateSequence() {
	// 	fmt.Printf("%d!! = %d - %d\n", 7, factorial.DoubleFactorial(7), s)
	// 	// break
	// }
	// fmt.Println(hashToSequence[targetToTreeHash[7]])

	// t := buildTree(945)
	// for s := range t.solver() {
	// 	fmt.Println(s)

	// }
}
