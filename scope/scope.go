package main

import (
	"fmt"
	"os"
)

var s = "I love Golang!"

func main() {
	fmt.Printf("%s\n", s) // print value of global s "I love Golang"

	var s = "I really love Go!" //new local variable
	fmt.Printf("%s\n", s)       //print "I really love Go!"

	//var s  = "I said you, I really love Go!" //its an error: 's' redeclared in this block
	//fmt.Printf("%s\n", s)
	fmt.Println("*****************")
	for i, s := range s {
		s := s // s belongs to this for loop, it's not for the visible upper scope
		fmt.Printf("%d:  %c\n", i, s)
	}
	fmt.Printf("%s\n", s)
	fmt.Println("*****************")

	length := len(s)
	fmt.Printf("%d\n", length)          //17
	if length := f(); length > length { //length is equal to 20 and 20>20 is not true, so the else block runs
		fmt.Printf("In if %d\n", length)
	} else {
		fmt.Printf("In else %d\n", length)
	}
	fmt.Printf("%d\n", length) // 17

	fmt.Println("*****************")

	err := g1()
	if err != nil {
		return
	}
}

func f() int {
	return 20
}

func g1() error {
	//if f, err := os.Open("scope.go"); err != nil { // compile error: unused: f
	//	return err
	//}
	//f.Stat() // compile error: undefined f because if is defined in this if block and not visible func block
	//f.Close()    // compile error: undefined f

	f, err := os.Open("scope.go") //no error because f is defined in this g1 func block
	if err != nil {
		return err
	}
	_, err = f.Stat()
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}
	return nil
}
