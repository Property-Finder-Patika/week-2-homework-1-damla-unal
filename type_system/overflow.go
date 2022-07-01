package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		width  uint8 = 255
		height       = 255 // int
	)

	fmt.Printf("width: %d\n", width)
	width++

	if int(width) < height {
		fmt.Println("height is greater")
	}

	fmt.Printf("width: %d height: %d\n", width, height)
	fmt.Println("***************************************")

	// go catches overflow at compile-time
	// fmt.Println(int8(math.MaxInt8 + 1)) // overflows: Cannot convert an expression of the type 'int' to the type 'int8'

	// but it cannot catch them in runtime
	var n int8 = math.MaxInt8

	// wrap around to its negative maximum
	fmt.Println("max int8     :", n)   // 127
	fmt.Println("max int8 + 1 :", n+1) // -128

	// wrap around to its maximum
	var un uint8

	fmt.Println("max uint8    :", un)   // 0
	fmt.Println("max uint8 - 1:", un-1) // 255

	// wrap around to its minimum
	un = math.MaxUint8
	fmt.Println("max uint8 + 1:", un+1) // 0

	// float goes to infinity when overflowed
	f := float32(math.MaxFloat32)
	fmt.Println("max float    :", f*1.2)

	fmt.Println("****************************")
	// uint16 max value is 65535
	big := uint16(65535)

	// uint8 destroys its value
	// to its own max value which is 255
	//
	// 65535 - 255 is lost.
	small := uint8(big)

	// fmt.Printf("%b %d\n", big, big)
	// fmt.Printf("%b %[1]d\n", big)

	fmt.Printf("%016b %[1]d\n", big)   //65535
	fmt.Printf("%016b %[1]d\n", small) //255
}
