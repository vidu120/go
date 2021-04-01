package main

import (
	"fmt"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func exists(node *TreeNode , storedHeights map[*TreeNode]int) bool{
	_ , exist := storedHeights[node]
	return exist
}

//recursive method to get the height of every node
func calHeightPerNode(root *TreeNode , storedHeights map[*TreeNode]int)int{
	if root == nil{
		return 0
	}
	if exists(root , storedHeights) {
		return storedHeights[root]
	}else{
		storedHeights[root] = 1  + int(math.Max(float64(calHeightPerNode(root.Left , storedHeights)) , float64(calHeightPerNode(root.Right ,storedHeights))))
	}
	return 0
}
//checks if the tree is balanced or not
func checkIfBalanced(root *TreeNode , storedHeights map[*TreeNode]int) bool{
	if root == nil{
		return true
	}

	fmt.Println(storedHeights[root.Left])
	fmt.Println(storedHeights[root.Right])

	if int(math.Abs(float64(storedHeights[root.Left]) - float64(storedHeights[root.Right]))) > 1{
		return false
	}
	return checkIfBalanced(root.Left , storedHeights) && checkIfBalanced(root.Right, storedHeights)
}

func isBalanced(root *TreeNode) bool {
	//we are mapping the TreeNode pointer to the specific height for that node
	storedHeights := make(map[*TreeNode]int)
	//we stored for every node
	_ = calHeightPerNode(root , storedHeights)

	//now check if it is balanced
	return checkIfBalanced(root , storedHeights)
}
