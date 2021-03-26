package main

import (
	"fmt"
	"log"
	"math"
)

func getIntInput() (int , error){
	var number int
	_ , err := fmt.Scan(&number)
	if err != nil{
		return 0, err
	}
	return number , nil
}


func checkPrime(number int) bool{
	sqrtNumber := math.Sqrt(float64(number))

	if number == 1{
		return false
	}

	for i := 2 ; i <= int(sqrtNumber);i++{
		if number % i == 0{
			return false
		}
	}
	return true
}
func main(){
	testCases , err := getIntInput()
	if err != nil{
		log.Fatal(err)
	}

	for testCases > 0{
		number , err := getIntInput()
		if err != nil{
			log.Fatal(err)
		}
		if checkPrime(number){
			fmt.Println("yes")
		}else{
			fmt.Println("no")
		}
		testCases--
	}
}


