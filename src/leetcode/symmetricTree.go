package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inOrderTraversal1(root *TreeNode , traversal []int) []int{
	if root == nil{
		return traversal
	}
	traversal = inOrderTraversal1(root.Left , traversal)
	traversal = append(traversal ,root.Val)
	traversal = inOrderTraversal1(root.Right , traversal)

	return traversal
}

func inOrderTraversal2(root *TreeNode , traversal []int) []int{
	if root == nil{
		return traversal
	}
	traversal = inOrderTraversal2(root.Right , traversal)
	traversal = append(traversal ,root.Val)
	traversal = inOrderTraversal2(root.Left , traversal)

	return traversal

}

func isSymmetricRecursive(leftTree , rightTree *TreeNode) bool{
	if leftTree == nil && rightTree == nil{
		return true
	}
	if (leftTree != nil && rightTree == nil) ||  (leftTree == nil && rightTree != nil) || (leftTree.Val != rightTree.Val) {
		return false
	}
	return isSymmetricRecursive(leftTree.Left , rightTree.Right) && isSymmetricRecursive(leftTree.Right , rightTree.Left)
}

//let's find the recursive solution for this approach
func isSymmetric(root *TreeNode) bool {
	return isSymmetricRecursive(root.Left , root.Right)
}