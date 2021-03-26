package datafile

import (
	"bufio"
	"os"
)

//GetStrings function gets the strings by parsing through the data file
func GetStrings(filename string) ([]string , error){

	var stringSlice []string

	//first of all get a pointer to the file
	file ,err := os.Open(filename)
	if err != nil{
		return nil ,err
	}

	//make a scanner variable
	scanner := bufio.NewScanner(file)

	//scan through the whole document
	for scanner.Scan(){
		stringSlice = append(stringSlice , scanner.Text())
	}

	//any error while closing the file
	err = file.Close()
	if err != nil{
		return nil, err
	}
	//any error while scanning through the file
	if scanner.Err() != nil{
		return nil, err
	}

	//return the slice of strings
	return stringSlice , nil
}
