package main

import (
	"fmt"
	"time"
)

func main() {
	const (
		a = 1
		b //1
		c = 2
		d //2
		e = 3
		f //3
	)
	fmt.Println(a, b, c, d, e, f)

	const i int = 5
	fmt.Printf("%d is %T\n", i, i)

	const j = 9 //removed int type
	fmt.Printf("%d is %T\n", j, j)

	const s2 = "heeyyooo"
	fmt.Printf("%s is %T\n", s2, s2)

	typelessConstant()
}

func typelessConstant() {
	// hours's type is time.Duration
	// but later's type was int
	// now, `later` is typeless
	//
	// time.Duration's underlying type is int64
	// and, `later` is numeric typeless value
	// so, they can be multiplied
	const later = 10
	fmt.Printf("%d is %T\n", later, later)

	hours, _ := time.ParseDuration("1h")

	fmt.Printf("%s later...\n", hours*later)

	printLast3Month()

	printCurrency()
}

func printLast3Month() {
	//initialize multiple constants using iota
	const (
		Dec = 12 - iota // 12 - 0 = 12
		Nov             // 12 - 1 = 11
		Oct             // 12 - 2 = 10
	)

	fmt.Println(Dec, Nov, Oct)
}

type Currency int

func printCurrency() {

	const (
		USD Currency = iota
		EUR
		GBP
		TL
		OTHER = -1
	)
	fmt.Println(USD, EUR, GBP, TL, OTHER)

}
