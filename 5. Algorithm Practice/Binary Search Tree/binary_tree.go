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
	arr1 := []int{4, 2, 7, 1, 3}
	l1 := arrayToTreeNode(arr1)
	print(l1, 0)

	//fmt.Println("l1 is ", isValidBST(l1))
	//fmt.Println("minimum of l1 is", minNode(l1).Val)

	// fmt.Println("inorder successor of 2 in l1 is ", inorderSuccessor(l1, 2))
	// fmt.Println(inorderTraversal(l1))

	print(insertIntoBST(l1, 5), 0)

}

func arrayToTreeNode(arr []int) *TreeNode {
	nodeArr := []*TreeNode{}
	n := len(arr)

	for _, v := range arr {
		if v == -1 {
			nodeArr = append(nodeArr, nil)
		} else {
			nodeArr = append(nodeArr,
				&TreeNode{
					Val:   v,
					Left:  nil,
					Right: nil,
				})
		}
	}

	for i, v := range nodeArr[0 : (n-1)/2+1] {
		if 2*i+1 < n {
			v.Left = nodeArr[2*i+1]
		}
		if 2*i+2 < n {
			v.Right = nodeArr[2*i+2]
		}
	}
	return nodeArr[0]

}

func print(root *TreeNode, n int) {
	if root != nil {

		space := ""
		for i := 0; i < n; i++ {
			space += "--"
		}
		if n > 0 {
			space = "|" + space + ">"
		}
		fmt.Println(space, root.Val)
		print(root.Left, n+1)
		print(root.Right, n+1)
	}
}

func inorderTraversal(tn *TreeNode) []int {
	result := []int{}

	if tn == nil {
		return result
	}

	if tn.Left != nil {
		result = append(result, inorderTraversal(tn.Left)...)
	}
	result = append(result, tn.Val)

	if tn.Right != nil {
		result = append(result, inorderTraversal(tn.Right)...)
	}

	return result
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

func isValidBST(root *TreeNode) bool {

	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}

	return helperIsValidBST(root, nil, nil)

}

func helperIsValidBST(root *TreeNode, minNode *TreeNode, maxNode *TreeNode) bool {
	if root == nil {
		return true
	}

	if (minNode != nil && root.Val <= minNode.Val) || (maxNode != nil && root.Val >= maxNode.Val) {
		return false
	}

	return helperIsValidBST(root.Left, minNode, root) && helperIsValidBST(root.Right, root, maxNode)

}

func inorderSuccessor(root *TreeNode, p int) *TreeNode {

	var inSucc *TreeNode
	for root != nil {

		if p < root.Val {
			inSucc = root
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return inSucc

}

/* Find the minimum */
func minNode(root *TreeNode) *TreeNode {
	min := root
	if root == nil {
		return nil
	}
	for min.Left != nil {
		min = min.Left
	}
	return min
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {

	if root == nil {
		return &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	}

	curr := root

	if val < curr.Val {
		curr.Left = insertIntoBST(curr.Left, val)
	} else {
		curr.Right = insertIntoBST(curr.Right, val)
	}

	return root

}

func missingNumber(nums []int) int {

	missing := len(nums)
	n := len(nums)

	if nums[0] < nums[n] {
		for i := 0; i < n; i++ {
			missing ^= nums[i] ^ i
		}

	} else {
		for i := 0; i < n; i++ {
			missing ^= nums[i] ^ (n - i)
		}
	}

	return missing

}

