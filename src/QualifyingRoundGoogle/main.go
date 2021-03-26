package main

import (
	"fmt"
	"log"
	"math"
)

//TakeSliceInput function for taking the input as string and then converting it to a slice int
func takeSliceInput(elems []int) []int {
	temp := 0
	for i := 0; i < cap(elems); i++ {
		_, err := fmt.Scan(&temp)
		if err != nil {
			log.Fatal(err)
		}
		elems = append(elems, temp)
	}
	return elems
}

func minElemIndex(elems []int) int {
	minElem := math.MaxInt32
	minIndex := 0
	for i := 0; i < len(elems); i++ {
		if elems[i] < minElem {
			minElem = elems[i]
			minIndex = i
		}
	}
	return minIndex
}

func reverseSort(elems []int) int {
	operationCost := 0
	for i := 0; i < len(elems)-1; i++ {
		j := i + minElemIndex(elems[i:])
		operationCost += reverseOperation(elems[i : j+1])
	}
	return operationCost
}

//this reverse operation will return the number of reverses performed
func reverseOperation(elems []int) int {

	for i := 0; i < len(elems)/2; i++ {
		temp := elems[i]
		elems[i] = elems[len(elems)-i-1]
		elems[len(elems)-i-1] = temp
	}

	return len(elems)
}

func main() {
	var testCases, numberOfElem int
	_, err := fmt.Scan(&testCases)
	if err != nil {
		log.Fatal(err)
	}
	for i := 1; i <= testCases; i++ {
		_, err = fmt.Scan(&numberOfElem)
		if err != nil {
			log.Fatal(err)
		}
		elems := make([]int, 0, numberOfElem)
		elems = takeSliceInput(elems)
		fmt.Printf("Case #%d: %d\n", i, reverseSort(elems))
	}
}
