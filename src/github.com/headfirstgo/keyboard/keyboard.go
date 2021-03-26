//Package keyboard gets the input from the user through the keyboard
package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)


//GetFloat function allows us to get the user input of float like values and also
//give of errors if any found while converting the user input given
func GetFloat()(float64 , error){
	//first make a reader to read from the standard input
	reader := bufio.NewReader(os.Stdin)

	//read the string from the stdin and also the error if any
	input , err := reader.ReadString('\n')
	if err != nil{
		return 0 , err
	}
	//trim the newlines , tabs from the beginning and the end
	input = strings.TrimSpace(input)

	//convert the string to the float type and check for any errors
	floatVariable , err := strconv.ParseFloat(input , 64)
	if err != nil{
		return 0 , err
	}

	return floatVariable ,nil
}