package main

import "fmt"

// run with this command : go run exercise_main.go exercise_printer.go
// If you just run go run exercise_main.go and that file has a reference to a function in another file within the same package,
// it will error because you didn't tell Go to run the whole package, you told it to only run that one file.

//You can tell go to run as a whole package by grouping the files as a package in the run command in several ways.
// like go run ./ or go run exercise_main.go exercise_printer.go
func main() {
	hello() //no need to import package because these are in same package
}

func bye() {
	fmt.Println("bye bye")
}
