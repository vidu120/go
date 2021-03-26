package main

import "strings"

//finds the first occurence of the substring needle in the haystack string
func strStr(haystack string, needle string) int {
	if len(haystack) == 0{
		return 0
	}
	index := strings.Index(haystack , needle)
	return index
}

func main(){
}