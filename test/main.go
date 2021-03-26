package main

import (
	"fmt"
	"github.com/headfirstgo/calendar"
	"log"
)

func main(){

	//this is the date today
	today := calendar.Date{}

	event := calendar.Event{}


	event.SetTitle("Program..")
	event.SetDate(today)


	err := today.SetYear(2012)
	if err != nil{
		log.Fatal(err)
	}
	err = today.SetMonth(12)
	if err != nil{
		log.Fatal(err)
	}
	err = today.SetDate(31)
	if err != nil {
		log.Fatal(err)
	}

	//get the values using the getter method
	fmt.Println(today.Year())
	fmt.Println(today.Month())
	fmt.Println(today.Date())
}
