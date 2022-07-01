package main

import "fmt"

func main() {
	// Uncomment the code line below to see the error.

	/*
		"Hello"
	*/

	// It will say: "evaluated but not used"
	//
	// Because:
	// "Hello" literal returns a value but there isn't any
	// statement which uses it.
	//
	// So:
	// You can't use expressions alone without statements.

	fmt.Println("*************************")
	// Operators combine multiple expressions together
	// as if there's a single expression.
	fmt.Println("Hello!" + "!" + "!" + "?")

	//print variable and type
	fmt.Printf("Type of %d is %[1]T\n", 3)

}
