package main

import "fmt"

func main() {
	// Variables are declared using the var keyword.
	// The type of the variable is declared after the variable name.
	// If the type is omitted, the type is inferred from the value.
	var a int
	var b = 2
	c := 3 // Short variable declaration

	// Variables can be declared and assigned in a single line.
	// Multiple variables can be declared and assigned in a single line.
	var d, e int = 4, 5
	f, g := 6, 7

	// Constants are declared using the const keyword.
	// Constants cannot be declared using the := syntax.
	const h = 8

	fmt.Println(a, b, c, d, e, f, g, h, i, j)
}

// Variables and constants can be declared at the package level.
var i int

const j = 9
