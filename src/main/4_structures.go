package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{
		Name: "John Doe",
		Age:  42,
	}

	fmt.Println(person)
}
