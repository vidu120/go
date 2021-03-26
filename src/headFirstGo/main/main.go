package main

import (
"fmt"
"reflect"
)

func main() {
	fmt.Println("Hey, their let's learn go")
	quantity := 10
	fmt.Println(quantity , reflect.TypeOf(quantity))
}