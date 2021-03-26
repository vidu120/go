package main

import (
	"errors"
	"fmt"
	"log"
)

func paintNeeded(width float64 , height float64) (float64 , error) {

	//make an error type variable
	var err error = nil

	if width < 0 && height < 0{
		err = fmt.Errorf("height %.2f and width %.2f can't be negative" , height , width)
		return 0 , err
	}else if height < 0{
		err = errors.New("height can't be negative")
		return 0 , err
	}else if width < 0{
		err = errors.New("width can't be negative")
		return 0 , err
	}

	area := width * height

	return area / 10.0 , err

}

func main(){

	paintNeed , err := paintNeeded(45 , 3.2)
	fmt.Println(err)
	fmt.Println(paintNeed)

	paintNeed , err = paintNeeded(-2.4 , 23)
	fmt.Println(err)
	fmt.Println(paintNeed)

	paintNeed , err = paintNeeded(23.4 , -9.9)
	fmt.Println(err)
	fmt.Println(paintNeed)

	paintNeed , err = paintNeeded(-9.3 , -4.32)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(paintNeed)
}
