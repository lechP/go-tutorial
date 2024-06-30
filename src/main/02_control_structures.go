package main

import "fmt"

func main() {

	// if-else statement
	x := 42
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than or equal to 10")
	}

	// switch statement
	y := 3
	switch y {
	case 1:
		fmt.Println("y is 1")
	case 2:
		fmt.Println("y is 2")
	default:
		fmt.Println("y is not 1 or 2")
	}

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
