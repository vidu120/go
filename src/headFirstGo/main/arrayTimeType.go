package main

import (
	"fmt"
	mathrand "math/rand"
	"reflect"
	"time"
)

func main(){

	var dates [3]time.Time

	generator  := mathrand.Int()
	mathrand.Seed(time.Now().Unix())

	//generated random number which will be different for each goroutine
	fmt.Println(generator)
	//type of the variable generator to generate random numbers
	fmt.Println(reflect.TypeOf(generator))

	dates[0] = time.Now()
	dates[1] = time.Unix(100000000000 , 0)
	dates[2] = time.Unix(dates[0].Unix() , 0)

	//print the dates that we got
	fmt.Printf("%#v" , dates)
}