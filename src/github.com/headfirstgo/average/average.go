//Package average calculated the average value from a given data file
package main

import (
	"fmt"
	"github.com/headfirstgo/datafile"
	"log"
)

//main function calculates the average of an integer array
func main(){
	//make an array
	numbers , err := datafile.GetFloats("data.txt")
	if err != nil{
		log.Fatal(err)
	}
	var sum float64
	for _ , value := range numbers{
		sum += value
	}
	sampleCount := float64(len(numbers))
	fmt.Println("Average :" , sum / sampleCount)
}