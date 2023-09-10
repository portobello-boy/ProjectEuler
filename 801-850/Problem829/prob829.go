package main

import (
	"fmt"
	"math"
	"numberTheory"
	"time"
)

var BOUND int64 = 31
var exploredTree map[int64]*Tree = make(map[int64]*Tree)
var leafCountRootMap map[int64]map[int64]struct{} = make(map[int64]map[int64]struct{})

type Tree struct {
	val   int64
	left  *Tree
	right *Tree
}

func newLeaf(value int64) *Tree {
	return &Tree{
		value,
		nil,
		nil,
	}
}

func buildTree(rootValue int64) *Tree {
	if val, found := exploredTree[rootValue]; found {
		return val
	}

	root := newLeaf(rootValue)
	left, right := optimalFactors(rootValue)
	if !(left == 1 && right == rootValue) {
		root.left = buildTree(left)
		root.right = buildTree(right)
	}
	leafCount := int64(root.countLeaves())
	exploredTree[rootValue] = root
	if _, ok := leafCountRootMap[leafCount]; !ok {
		leafCountRootMap[leafCount] = make(map[int64]struct{})
	}
	leafCountRootMap[leafCount][root.val] = struct{}{} // map count of leaves to "set" of root values
	return root
}

func optimalFactors(value int64) (int64, int64) {
	sqrt := int64(math.Sqrt(float64(value)))
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

func (t *Tree) countLeaves() int {
	if t.isLeaf() {
		return 1
	}
	return t.left.countLeaves() + t.right.countLeaves()
}

func (t *Tree) isSameStructure(other *Tree) bool {
	if other == nil {
		return false
	}
	tLeaf := t.isLeaf()
	oLeaf := other.isLeaf()

	if tLeaf != oLeaf {
		return false
	} else if tLeaf && oLeaf {
		return true
	}
	return t.left.isSameStructure(other.left) && t.right.isSameStructure(other.right)
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

func (t *Tree) setLeavesToTwo() {
	if t.isLeaf() {
		t.val = 2
	} else {
		t.left.setLeavesToTwo()
		t.right.setLeavesToTwo()
		t.val = t.left.val * t.right.val
	}
}

func (t *Tree) buildMatchingTreeRoot() int64 {
	// Recurse depth-first left, build matching subtree where all leaves are 2
	candidateTree := t.copy()
	candidateTree.left.setLeavesToTwo()
	candidateTree.val = candidateTree.left.val * candidateTree.right.val

	// Find first right subtree which matches the base right subtree
	candidateRight := t.right.val
	for rootVal := range leafCountRootMap[int64(candidateTree.right.countLeaves())] {
		if rootVal < candidateTree.left.val {
			continue
		}
		if rootVal < candidateRight && candidateTree.right.isSameStructure(buildTree(rootVal)) {
			candidateRight = rootVal
		}
	}

	fmt.Println("Checking range:", candidateTree.left.val, candidateRight)

	for i := candidateTree.left.val; i <= candidateRight; i++ {
		iTree := buildTree(i)

		if !iTree.isSameStructure(candidateTree.right) {
			continue
		}
		candidateRight = i
		break
	}
	// return product of left and right root vals
	return candidateTree.left.val * candidateRight
}

func M(n int64) int64 {
	nTree := buildTree(numberTheory.DoubleFactorial(int(n)))
	if nTree.left == nil {
		return 2
	}

	return nTree.buildMatchingTreeRoot()

	// var leftSubtree *Tree = nil
	// var rightSubtree *Tree = nil

	// for i := int64(2); ; i++ {
	// 	if i%100000 == 0 {
	// 		fmt.Println(n, i)
	// 	}
	// 	iTree := buildTree(i)

	// 	if !iTree.isSameStructure(nTree.left) {
	// 		continue
	// 	}

	// 	leftSubtree = iTree
	// 	break
	// }

	// for i := leftSubtree.val; ; i++ {
	// 	if i%100000 == 0 {
	// 		fmt.Println(n, i)
	// 	}
	// 	iTree := buildTree(i)

	// 	if !iTree.isSameStructure(nTree.right) {
	// 		continue
	// 	}
	// 	rightSubtree = iTree
	// 	break
	// }

	// treeCandidate := &Tree{
	// 	leftSubtree.val * rightSubtree.val,
	// 	leftSubtree,
	// 	rightSubtree,
	// }

	// if !nTree.isSameStructure(treeCandidate) {
	// 	panic(errors.New("fuck"))
	// }

	// if n == 16 {
	// 	fmt.Println("waddup")
	// }

	// return leftSubtree.val * rightSubtree.val

	/* */

	// for i := 2; ; i++ {
	// 	if i%100000 == 0 {
	// 		fmt.Println(n, i)
	// 	}
	// 	iTree := buildTree(i)
	// 	if nTree.isSameStructure(iTree) {
	// 		return i
	// 	}
	// }
}

func main() {
	// t := buildTree(numberTheory.DoubleFactorial(9))
	// u := t.copy()
	// t.Print()
	// u.Print()
	// fmt.Println(t.buildMatchingTreeRoot())

	for i := int64(0); i < int64(math.Floor(math.Sqrt(float64(numberTheory.DoubleFactorial(int(BOUND)))))); i++ {
		_ = buildTree(i)
	}

	sum := int64(0)
	for n := int64(2); n <= BOUND; n++ {
		t := time.Now()
		Mn := M(n)
		sum += Mn
		fmt.Println(n, Mn, time.Since(t))
	}
	fmt.Println(sum)
}
