package main

import (
	"strconv"
	"unicode/utf8"
)

//this gets the int format of the string that we have got
func getValueFromChar(a string) int {
	val, _ := strconv.Atoi(a)
	return val
}

//function to implement
func addBinary(a string, b string) string {

	//our final string
	newString := ""

	sum := 0
	carry := 0

	//two counters
	counter1 := utf8.RuneCountInString(a) - 1
	counter2 := utf8.RuneCountInString(b) - 1

	for counter1 >= 0 && counter2 >= 0 {
		sum = getValueFromChar(string(a[counter1])) + getValueFromChar(string(b[counter2])) + carry
		newString = strconv.Itoa(sum%2) + newString
		carry = sum / 2
		//decrement the counters for both the string
		counter2--
		counter1--
	}
	for counter1 >= 0 {
		sum = getValueFromChar(string(a[counter1])) + carry
		newString = strconv.Itoa(sum%2) + newString
		carry = sum / 2
		counter1--
	}

	for counter2 >= 0 {
		sum = getValueFromChar(string(b[counter2])) + carry
		newString = strconv.Itoa(sum%2) + newString
		carry = sum / 2
		counter2--
	}
	if carry > 0 {
		newString = strconv.Itoa(carry%2) + newString
	}
	return newString
}

//
//func main(){
//	first := "111"
//	second  := "10101"
//	fmt.Println(addBinary(first , second))
//}
