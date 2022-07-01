package main

import "fmt"

// Package initialization begins by initializing package-level variables in the order in which they are declared,
// except that dependencies are resolved first
var j = i + 1 //j initialized second : 6

func createB() bool { //second createB runs and return true
	fmt.Println("In CreateB")
	return true
}

var i = createI() //i initialized first: 5

func main() {
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(b)
}

func createI() int { //first createI runs and return 5
	fmt.Println("In createI")
	//fmt.Printf("%d %d \n", i, j) //initialization loop error, j refers to i, i refers to createI, createI refers to j
	return 5
}

//init functions can’t be called or referenced, but otherwise they are normal functions.
//Within each file, init functions are automatically executed when the program starts,
//in the order in which they are declared
func init() {
	fmt.Println("In init I")
	fmt.Printf("%d %d \n", i, j) // 5,6
}

func init() {
	fmt.Println("In init II")
	fmt.Printf("%d %d \n", i, j) // 5,6
}

var b = createB() //b initialized third: true
