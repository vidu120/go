package main

import (
	"strings"
)

//get the length of the last word
func lengthOfLastWord(s string) int {

	//trimmed string
	s = strings.TrimSpace(s)
	lengthOfWord := 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			break
		}
		lengthOfWord++
	}

	return lengthOfWord
}

//
//func main(){
//	str := "a "
//	fmt.Println(lengthOfLastWord(str))
//	str = " "
//	fmt.Println(lengthOfLastWord(str))
//}
