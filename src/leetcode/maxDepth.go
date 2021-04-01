package main

//function to implement

//tested this function and this doesn't give the best results for our question
//func calDepth(root *TreeNode) int {
//	if root == nil{
//		return 0
//	}
//	return int(math.Max(float64(1 + maxDepth(root.Left)) , float64(1 + maxDepth(root.Right))))
//}


//function to implement
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return left + 1
	}
	return right+1
}