package main

import "fmt"

//implement this function
func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums);i++{
		if nums[i] == val{
			nums = append(nums[:i] , nums[i + 1:]...)
			i--
		}
	}
	return len(nums)
}


func main(){
	//make a slice literal

	mySlice := []int{2 , 4 , 5, 4 , 55 , 5 ,5}
	fmt.Println("This is the length of the slice" , len(mySlice))
	length := removeElement(mySlice , 4)
	fmt.Println("This is the length of the slice after removing the element"  , length)
	fmt.Println(mySlice)
}