package main

import (
	"fmt"
	"math"
)

//function to implement

//the below function is the simple O(n) way
//func maxSubArray(nums []int) int {
//
//	//some variables that are required
//	maxTillHere := 0
//	maxSub := math.MinInt32
//
//	//check for the max contiguous subarray in one parse
//	for _ , val := range nums{
//		maxTillHere += val
//		if maxSub < maxTillHere{
//			maxSub = maxTillHere
//		}
//		if maxTillHere < 0{
//			maxTillHere = 0
//		}
//	}
//	return maxSub
//}

//using divide and conquer method
func maxSubArray(nums []int) int {
	return divideAndConquer(nums)
}

//the divide and conquer approach
func divideAndConquer(nums []int) int{
	if len(nums) == 0{
		return math.MinInt32
	}
	if len(nums) == 1{
		return nums[0]
	}
	mid := len(nums) / 2
	return int(math.Max(float64(divideAndConquer(nums[:mid])) , math.Max(float64(divideAndConquer(nums[mid:])) , float64(combined(nums)))))
}

//return the combined max result
func combined(nums []int) int{
	rightMax := 0
	leftMax := 0

	start := 0

	if len(nums) % 2 == 0{
		leftMax = nums[len(nums) / 2 - 1]
		rightMax = nums[len(nums) / 2]
		start = len(nums) / 2 - 1
	}else{
		leftMax = nums[len(nums) / 2]
		rightMax = nums[len(nums) / 2 + 1]
		start = len(nums) / 2
	}

	sum := 0
	for i := start; i >= 0 ; i--{
		sum += nums[i]
		leftMax = int(math.Max(float64(leftMax) ,float64(sum)))
	}
	sum = 0
	for i := start + 1; i < len(nums) ; i++{
		sum += nums[i]
		rightMax = int(math.Max(float64(rightMax) , float64(sum)))
	}

	if leftMax < 0 || rightMax < 0{
		return int(math.Max(float64(leftMax) , float64(rightMax)))
	}
	return leftMax + rightMax
}

func main(){
	//our test function
	nums := []int{2 , 4  , -3 , -2 , 21  , -11}
	fmt.Println(maxSubArray(nums))
}