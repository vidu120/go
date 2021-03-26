//pass_fail go file determines whether a student is passing or failing
package main

import (
	"fmt"
	"github.com/headfirstgo/keyboard"
	"log"
)

//main This basically tells whether a person is passing or failing
func main(){

	//prompt the user to input the grade
	fmt.Print("Enter a grade : ")
	grade , err := keyboard.GetFloat()

	if err != nil{
		log.Fatal(err)
	}

	status := ""
	//check if the person passed or not
	if grade >= 60{
		status = "passing"
	}else{
		status = "failing"
	}

	//printing the input back to the user
	fmt.Println("The person is" , status)
}