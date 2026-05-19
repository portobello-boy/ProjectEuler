package main

import (
	"MathLibGo/numberTheory/primes"
	"MathLibGo/operations/factorial"
	"fmt"
	"math"
	"time"
)

var BOUND uint64 = 31
var hashMapToValue = map[uint64][]uint64{}
var hashMapIndex = uint64(2)

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
	left, right := optimalFactors(rootValue)
	if !(left == 1 && right == rootValue) {
		root.left = buildTree(left)
		root.right = buildTree(right)
	}
	return root
}

func optimalFactors(value uint64) (uint64, uint64) {
	sqrt := uint64(math.Sqrt(float64(value)))
	for n := sqrt; n > 1; n-- {
		if value%n == 0 {
			return n, value / n
		}
	}
	return 1, value
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
	frontier := []*Tree{t}
	hash := uint64(0)

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

	return hash
}

func getMinimalTreeValue(hash uint64, lowerBound uint64) uint64 {
	if values, ok := hashMapToValue[hash]; ok {
		for _, value := range values {
			if value > lowerBound {
				return value
			}
		}
	}

	t := buildTree(uint64(hashMapIndex))

	for t.hash() != hash || hashMapIndex < lowerBound {
		if _, ok := hashMapToValue[t.hash()]; !ok {
			hashMapToValue[t.hash()] = make([]uint64, 0)
		}
		hashMapToValue[t.hash()] = append(hashMapToValue[t.hash()], uint64(hashMapIndex))
		hashMapIndex++
		t = buildTree(uint64(hashMapIndex))
	}

	return hashMapIndex
}

func S(n *Tree, lowerBound uint64) uint64 {
	if n.isLeaf() {
		p := 2
		for p < int(lowerBound) {
			p, _ = primes.NextPrime(int(lowerBound))
		}
		return uint64(p)
	}

	nHash := n.hash()

	if values, ok := hashMapToValue[nHash]; ok {
		for _, value := range values {
			if value > lowerBound {
				return value
			}
		}
	}

	// Find the lowest value matching the hash of the left child
	lv := S(n.left, 0)

	// Find the lowest value matching the hash of the right child, value must be greater than lv
	rv := S(n.right, lv)

	val := lv * rv

	if lv*rv < lowerBound {
		val = getMinimalTreeValue(nHash, lowerBound)
	}

	// Add to map
	if _, ok := hashMapToValue[nHash]; !ok {
		hashMapToValue[nHash] = make([]uint64, 0)
	}
	hashMapToValue[nHash] = append(hashMapToValue[nHash], lv*rv)

	return val
}

func M(n uint64) uint64 {
	return S(buildTree(factorial.DoubleFactorial(uint64(n))), 0)
}

func main() {
	// fmt.Println(S(buildTree(945), 0))

	// t := buildTree(factorial.DoubleFactorial(9))
	// // u := t.copy()
	// t.Print()
	// fmt.Printf("hash: %b\n", t.hash())
	// t = buildTree(72)
	// t.Print()
	// fmt.Printf("hash: %b\n", t.hash())
	// t = buildTree(3840)
	// t.Print()
	// fmt.Printf("hash: %b\n", t.hash())

	fmt.Println(M(5))
	fmt.Println(M(7))
	fmt.Println(M(8))
	fmt.Println(M(12))

	// t = buildTree(180)
	// fmt.Printf("hash: %b\n", t.hash())

	// u.Print()
	// fmt.Println(t.buildMatchingTreeRoot())

	// for i := int64(0); i < int64(math.Floor(math.Sqrt(float64(factorial.DoubleFactorial(int64(BOUND)))))); i++ {
	// 	_ = buildTree(i)
	// 	if i%10 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	sum := uint64(0)
	for n := uint64(2); n <= BOUND; n++ {
		t := time.Now()
		Mn := M(n)
		sum += Mn
		fmt.Println(n, Mn, time.Since(t))
	}
	fmt.Println(sum)
}
