package main

import (
	"fmt"
	"strconv"
)

func main() {

	convertStrToInt()
	convertIntToStr()
}

func convertStrToInt() {
	fmt.Println("String to Integer")
	i, _ := strconv.Atoi("1291")
	fmt.Println(i)
	i++
	fmt.Println(i)
}

func convertIntToStr() {
	fmt.Println("Integer To String")
	i := 12
	stringI1 := fmt.Sprintf("%d", i)
	fmt.Println(stringI1)
	stringI1 += ":"
	fmt.Println(stringI1)

	strI := strconv.Itoa(i) // int to string
	fmt.Printf("strI: %s and type is %T\n", strI, strI)
}
