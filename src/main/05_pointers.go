package main

import "fmt"

func main() {
	a := 42
	b := &a

	fmt.Println(a, *b)
	*b = 21
	// 'a' is now 21 as well
	// 'b = 21' won't compile as b is a "pointer" - b is of type *int which means it can only store the memory address of an int variable (and isn't an int variable itself)
	fmt.Println(a, *b)

	// memory address of 'a' and the value of 'b' (which is the memory address of 'a') are the same
	fmt.Println(&a, b)
}
