package main

import (
	"errors"
	"fmt"
)

func returnError(a, b int) error {
	if a == b {
		err := errors.New("Error in returnError() function!")
		return err
	} else {
		return nil
	}
}

func main() {
	err := returnError(1, 2)
	if err == nil {
		fmt.Println("returnError() ended normally!")
	} else {
		fmt.Println(err)
	}

	err = returnError(10, 10)
	if err == nil {
		fmt.Println("returnError() ended normally!")
	} else {
		fmt.Println(err)
	}

	// Allows me to convert an error variable into a string variable
	if "Error in returnError() function!" == err.Error() {
		fmt.Println("!!")
	}
}
