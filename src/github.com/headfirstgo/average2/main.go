//Package computes the average of the command line arguments
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

//Average function returns the average of an array of float type
func average(numbers ...float64) float64{
	sum := 0.0
	for _ , number := range numbers{
		sum += number
	}
	return sum / float64(len(numbers))
}


func main(){

	//get the command line arguments that we have passed to the program
	args := os.Args[1:]
	var numbers []float64

	for _, arg := range args{
		temp ,err := strconv.ParseFloat(arg , 64)
		if err != nil{
			log.Fatal(err)
		}
		numbers = append(numbers ,temp)
	}

	fmt.Println("Average2 :" , average(numbers...))
}