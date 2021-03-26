package main

import "fmt"

//main function helps in understanding the working of slices
func main(){

	s1 := []int{1 , 2 , 3}
	fmt.Println(cap(s1))
	s2 := append(s1 , 4 , 5)
	fmt.Println(cap(s2))
	s3 := append(s2 , 6 , 7)
	fmt.Println(cap(s3))
	s4 := append(s3 , 8 , 9)
	fmt.Println(cap(s4))
	fmt.Println(s1 , s2 , s3 , s4)
	s4[0] = 0
	fmt.Println(s1 , s2 , s3 , s4)
}
