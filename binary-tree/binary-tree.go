package main

import "fmt"

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

}