//guess challenges players to guess a random number
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	//get the seconds right now
	seconds := time.Now().Unix()

	//seed the random generator to generate values every time
	rand.Seed(seconds)

	//make a random number between 1 and 100
	target := rand.Intn(100) + 1

	fmt.Println("I've chosen a random number between 1 and 100")
	fmt.Println("Can you guess the number ??")

	//fmt.Println(target)

	//make a reader variable
	reader := bufio.NewReader(os.Stdin)

	success := false

	for guesses := 0 ; guesses < 10;guesses++{
		fmt.Println("You have", 10 - guesses , "guesses left")
		fmt.Print("Make a guess - ")
		input , err := reader.ReadString('\n')
		if err != nil{
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess , err := strconv.Atoi(input)
		if err != nil{
			log.Fatal(err)
		}
		if guess < target{
			fmt.Println("Guess is too low")
		}else if guess > target{
			fmt.Println("Guess is too high")
		}else{
			fmt.Println("You got the guess right..!!")
			success = true
			break
		}
	}

	fmt.Printf("hey  ,there , the target was %d"  ,target)
	store := fmt.Sprintf("hey , there the targer was %d" , target)
	fmt.Println(store)

	if !success{
		fmt.Println("Sorry , you didn't guess the number right ,It was :" , target)
	}
}
