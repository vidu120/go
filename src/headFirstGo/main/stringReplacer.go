package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main(){

	myString :="Hell# , h#w are y#u."

	//var replacer strings.Replacer = strings.NewReplacer("#")
	replacer := strings.NewReplacer("#" , "o")

	var mine *strings.Replacer = strings.NewReplacer("cidhn" , "vidhan")

	fmt.Println(reflect.TypeOf(mine))
	fmt.Println(reflect.TypeOf(replacer))

	fmt.Println("Originl String - ", myString)
	fmt.Println("Fixed String - " , replacer.Replace(myString))
}
