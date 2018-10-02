package main

import "fmt"

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

	fmt.Println(inorderTraversal(&tn1))

}

func inorderTraversal(tn *TreeNode) []int {
	result := []int{}

	if tn == nil {
		return result
	}

	if tn.Left != nil {
		result = append(result, inorderTraversal(tn.Left)...)
	}

	if tn.Right != nil {
		result = append(result, inorderTraversal(tn.Right)...)
	}

	result = append(result, tn.Val)

	return result
}
