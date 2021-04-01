package main

import "fmt"

func main(){
	firstArray := []int{1 , 2 , 4 , 5 , 0 , 0  , 0}
	secondArray := []int{6 , 6 , 6}
	fmt.Println("Hey , how are you")
	merge(firstArray ,4 , secondArray , 3)
	fmt.Println(firstArray)
}