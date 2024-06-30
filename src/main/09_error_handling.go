package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("can't divide by zero")
	}
	return a / b, nil
}

func main() {
	result, err := divide(1, 0)
	if err != nil {
		fmt.Println("Error: ", err)
		// panic(err) - ungraceful exit
	} else {
		fmt.Println("Result: ", result)
	}
}
