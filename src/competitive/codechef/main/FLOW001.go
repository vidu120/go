package main

import (
	"fmt"
	"log"
)


func main(){
	testCases , err := getIntInput()
	if err != nil{
		log.Fatal(err)
	}

	for testCases > 0{
		number1 , err := getIntInput()
		if err != nil{
			log.Fatal(err)
		}
		number2 , err := getIntInput()
		if err != nil{
			log.Fatal(err)
		}
		fmt.Println(number1 + number2)
		testCases--
	}
}


