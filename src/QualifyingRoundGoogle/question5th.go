package main

import (
	"fmt"
	"log"
)

func takeInput(elems []string) {
	temp := ""
	for i := 0; i < len(elems); i++ {
		_, err := fmt.Scan(&temp)
		if err != nil {
			log.Fatal(err)
		}
		elems[i] = temp
	}
}

func main() {
	testCases := 0
	quizPercent := 0
	_, err := fmt.Scan(&testCases)
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Scan(&quizPercent)
	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i <= testCases; i++ {
		var resultPerPerson [100]string
		takeInput(resultPerPerson[:])
		fmt.Println(resultPerPerson)
	}
}
