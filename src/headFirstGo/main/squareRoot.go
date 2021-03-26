package main

import (
	"fmt"
	"log"
	"math"
)

//we can even specify the name of the return values so that it can be of help to the developer

//also try to make the function in such a way so that it returns an error bundled with our required values
//This in turn can help in the long run when any of our processes get exhausted out
func getSquareRoot(number int) (squareRoot float64 , err error){
	if number < 0 {
		return 0, fmt.Errorf("cannot find a square root of a negative number")
	}
	return math.Sqrt(float64(number)) , nil
}

func main(){
	squareRoot , err := getSquareRoot(-90)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(squareRoot)
}