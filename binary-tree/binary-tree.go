package main

import (
	"fmt"
	"math"
)

//node type
type BinaryNode struct {
	data  int64
    left  *BinaryNode
    right *BinaryNode
    
}

type BinaryTree struct {
	root *BinaryNode
}

func depthFirstValues(bt BinaryTree) []int64 {
	if bt.root == nil {
		return []int64{}
	}
	stack := []*BinaryNode{bt.root}
	result := []int64{}
	for len(stack) != 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, current.data)
		if current.left != nil {
			stack = append(stack, current.left)
		}

		if current.right != nil {
			stack = append(stack, current.right)
		}
	}

	return result
}

func depthFirstValuesRecursive(root *BinaryNode) []int64 {
	if root == nil {
		return []int64{}
	}

	leftVal := depthFirstValuesRecursive(root.left)
	rightVal := depthFirstValuesRecursive(root.right)

	return append([]int64{root.data}, append(leftVal, rightVal...)...)
}

func breadthFirstValues(bt BinaryTree) []int64 {
	if bt.root == nil {
		return []int64{}
	}

	values := []int64{}
	queue := []*BinaryNode{bt.root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		values = append(values, current.data)

		if current.left != nil {
			queue = append(queue, current.left)
		}

		if current.right != nil {
			queue = append(queue, current.right)
		}
	}

	return  values
}

//find a value
func treeIncludes(root *BinaryNode, target int64) bool {
	if root == nil {
		return false
	}

	queue := []*BinaryNode{root}

	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]

		if current.data == target {
			return true
		}

		if current.left != nil {
			queue = append(queue, current.left)
		}

		if current.right != nil {
			queue = append(queue, current.right)
		}
	}

	return false
}


func treeIncludesRecursive(root *BinaryNode, target int64) bool {
	if root == nil {
		return false
	}

	if root.data == target {
		return true
	}
	left := treeIncludes(root.left, target)
	right :=treeIncludes(root.right, target)

	return left || right
}

//tree sum
//tree sum recursive
func treeSumRecursive(root *BinaryNode) int64 {
	if root == nil {
		return 0
	}

	return root.data + treeSumRecursive(root.left) + treeSumRecursive(root.right)
}

func treeSum(root *BinaryNode) int64 {
	if root == nil {
		return 0
	}
	var sum int64 = 0
	queue := []*BinaryNode{root}

	for len(queue) != 0 {
		current := queue[0]
		sum += current.data
		queue = queue[1:]

		if current.right != nil {
			queue = append(queue, current.right)
		}

		if current.left != nil {
			queue = append(queue, current.left)
		}
	}

	return sum
}

func treeMin(root *BinaryNode) int64 {
	smallest := int64(math.MaxInt64)
	if root == nil {
		return smallest
	}

	stack := []*BinaryNode{root}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current.data < int64(smallest) {
			smallest = current.data
		}

		if current.right != nil {
			stack = append(stack, current.right)
		}

		if current.left != nil {
			stack = append(stack, current.left)
		}
	}

	return int64(smallest)

}

func treeMinRecursive(root *BinaryNode) int64 {
	if root == nil {
		return math.MaxInt64
	}

	leftMin := treeMinRecursive(root.left)
	rightMin := treeMinRecursive(root.right)

	return min(root.data, leftMin, rightMin)
}

func maxRootToLeafSum(root *BinaryNode) int64 {
	if root == nil {
		return math.MinInt64
	}

	if root.left == nil && root.right == nil {
		return root.data
	}

	maxChildPathSum := math.Max(float64(maxRootToLeafSum(root.left)), float64(maxRootToLeafSum(root.right)))

	return root.data + int64(maxChildPathSum)
}

func main() {
	//define nodes
	node1 := &BinaryNode{data: 1}
	node2 := &BinaryNode{data: 2}
	node3 := &BinaryNode{data: 3}
	node4 := &BinaryNode{data: 4}
	node5 := &BinaryNode{data: 5}
	node6 := &BinaryNode{data: 6}

	node1.left = node2
	node1.right = node3
	node2.left = node4
	node2.right = node5
	node3.right = node6

	tree := BinaryTree{root: node1}
	fmt.Println("DEPTH FIRST VAL")
	fmt.Println(depthFirstValues(tree))
	fmt.Println(depthFirstValuesRecursive(node1))
	fmt.Println("BREADTH FIRST VAL")
	fmt.Println(breadthFirstValues(tree))
	fmt.Println("FIND VALUE ")
	fmt.Println("FIND VALUE USING NON-RECURSIVE BFS: ", treeIncludes(node1, 7))
	fmt.Println("FIND VALUE USING RECURSIVE BFS: ", treeIncludesRecursive(node1, 6))
	fmt.Println("TREE SUM RECURSIVE: ", treeSumRecursive(node1))
	fmt.Println("TREE SUM: ", treeSum(node1))
	fmt.Println("TREE MIN VALUE: ", treeMin(node1))
	fmt.Println("TREE MIN VALUE RECURSIVE DFS: ", treeMinRecursive(node1))
	fmt.Println("MAX ROOT TO LEAF VALUE: ", maxRootToLeafSum(node1))
}