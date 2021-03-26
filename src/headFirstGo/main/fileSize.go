//getting the size of a file
package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
)

func main(){

	//gets the file info pointer variable to the original variable created
	fileInfo , err := os.Stat("")
	fmt.Println(reflect.TypeOf(fileInfo))

	//captures the error and displays it
	if err != nil {
		log.Fatal(err)
	}

	//prints the size of the file
	fmt.Println(fileInfo.Size())

	fmt.Println(reflect.TypeOf(fileInfo))
	fmt.Println(reflect.TypeOf(err))

}
