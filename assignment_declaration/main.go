package main

import (
	"fmt"
	"math"
)

//s1 := "An example declaration." // .= can only be used for local variables
var s11 string = "An example declaration."
var s22 = "Another declaration."
var s33 string

func main() {
	fmt.Println("-------------Assignment examples------------")
	x := 1 // named variable

	b := false // named variable
	p := &b    // named variable, p points to address of b
	*p = true  // indirect variable

	person := Person{}  // named variable
	person.name = "bob" // struct field

	fmt.Printf("%d %t %s\n", x, *p, person.name)

	y := 9
	x, y = y, x //firstly, calculate right side then assign to left
	fmt.Printf("%d %d\n", x, y)

	fmt.Println("-------------Conversion examples------------")
	fp1 := 3.14
	i := int(fp1) // 3
	fmt.Println(i)

	//i = int(6.28) // Cannot convert an expression of the type 'float64' to the type 'int'
	i = int(math.Floor(6.28))
	fmt.Println(i)

	type yazi string
	var d yazi = "I like Go!"

	type text string
	var t text = "I like Go!"

	//fmt.Println(d == t) // Invalid operation: d == t (mismatched types yazi and text)

	var s1 = string(d)
	var s2 = string(t)
	fmt.Println(s1, s2)
	fmt.Println(s1 == s2) //true

	fmt.Println("-------------Declaration examples------------")
	// Shorter way, type is omitted, the type is inferred from the value
	var anotherInteger = 145
	fmt.Println("Another integer: ", anotherInteger)

	i1, j := 3, 5 // both of them new declaration
	fmt.Printf("%d %d\n", i1, j)

	i1, j, k := 3, 5, 7.5 // j and k are new declaration
	fmt.Printf("%d %d %f\n", i1, j, k)

	//i1, j, k := 13, 15, 17 // Compile time error due to no new declaration

	var isLiquid bool //declaration
	_ = isLiquid      // discard it using a blank-identifier to be unused variable
}

type Person struct {
	name string
}
