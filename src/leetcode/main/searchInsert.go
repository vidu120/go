package main

import "fmt"

//function to implement
func searchInsert(nums []int, target int) int {
	//first let's apply binary search to find the index of the target
	result := binarySearch(nums , target)
	if result != -1{
		return result
	}
	return binaryInsertionIndex(nums, target)
}

//find the position to insert the target point if it is not found in the original array
func binaryInsertionIndex(nums []int , target int) int{
	start := 0
	end := len(nums) - 1
	for start < end{
		mid := start + (end - start) / 2
		if nums[mid] < target{
			start = mid + 1
		}else{
			end = mid - 1
		}
	}
	if nums[start] < target{
		return start + 1
	}
	return start
}

//binary search implementation
func binarySearch(nums []int ,target int) int{
	//let's use an iterative version
	start := 0
	end := len(nums) - 1
	for start <= end{
		mid := start + (end - start) / 2
		if nums[mid] == target{
			return mid
		}else if nums[mid] < target{
			start = mid + 1
		}else{
			end = mid - 1
		}
	}
	return -1
}

func main(){

	//let's take some test here
	arr := []int{1 ,3 , 5 , 6}
	fmt.Println(searchInsert(arr , 3))
	fmt.Println(searchInsert(arr, 4))
	fmt.Println(searchInsert(arr , 5))
}