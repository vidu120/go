package leetcode

import (
	"fmt"
	"github.com/headfirstgo/calendar"
	"log"
)

func main() {
	event := calendar.Event{}
	err := event.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDate(30)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetYear(2000)
	if err != nil {
		log.Fatal(err)
	}
	//get the values using the getter method
	fmt.Println(event.Year())
	fmt.Println(event.Month())
}
