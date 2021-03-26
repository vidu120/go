package main

import "fmt"

func main(){
	fmt.Println(sum(4 , 1 , 9 , 2))
	fmt.Println(sum(7))
}

//Sum returns the sum of the numbers given
func sum(numbers ...int) int{
	var sum = 0
	for _ , number := range numbers{
		sum += number
	}
	return sum
}