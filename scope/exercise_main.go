package main

import "fmt"

func main() {
	hello() //no need to import package because these are in same package
}

func bye() {
	fmt.Println("bye bye")
}
