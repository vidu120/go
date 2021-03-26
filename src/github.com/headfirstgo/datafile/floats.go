//Package datafile allows reading data files from input
package datafile

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//GetFloats function reads and stores float value line by line
func GetFloats(filename string) (floats []float64 , err error){

	//make a 3 numbers array
	var numbers []float64

	//get the pointer to the file
	file ,err := os.Open(filename)
	if err != nil{
		return nil , err
	}

	scanner := bufio.NewScanner(file)
	//make a temp variable
	number := 0.0
	//scanning the text per line and trimming and converting it to a float value
	for scanner.Scan() {
		number, err = strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers , number)
	}
	//close the file so as to free the computer resources
	err = file.Close()
	if err != nil{
		return nil ,err
	}
	//if we encountered any scanning problem while scanning through the document
	if scanner.Err() != nil{
		return nil , scanner.Err()
	}
	//return the perfect answer
	return numbers , nil
}
