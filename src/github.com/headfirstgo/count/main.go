package main

import (
	"fmt"
	"github.com/headfirstgo/datafile"
	"log"
)

func main() {

	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)
	for _ ,name := range lines{
		counts[name]++
	}

	//print out the name and their counts to the console
	for name , count := range counts{
		fmt.Printf("Votes for %s: %d\n" , name , count)
	}

}
