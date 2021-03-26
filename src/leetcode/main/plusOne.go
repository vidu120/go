package main

func plusOne(digits []int) []int {
	index := len(digits) - 1
	for index >= 0 && digits[index] == 9{
		digits[index] = 0
		index--
	}

	if index < 0{
		//so basically if the index is negative then we need to append 1 to the head of the integer array
		digits = append(digits , 0)
		digits = append(digits[1:] , digits[:len(digits) - 1]...)
		digits[0] = 1
		return digits
	}
	digits[index]++
	return digits
}


func main(){

}