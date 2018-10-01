package main

import (
	"fmt"
)

//TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	var tn1, tn2, tn3 TreeNode
	tn1.Val = 1
	tn1.Left = &tn2
	tn1.Right = &tn3
	tn2.Val = 2
	tn3.Val = 3

	fmt.Println(preorderTraversal(&tn1))

}

func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}

	result = append(result, root.Val)

	if root.Left != nil {
		result = append(result, preorderTraversal(root.Left)...)
	}

	if root.Right != nil {
		result = append(result, preorderTraversal(root.Right)...)
	}

	return result
}
