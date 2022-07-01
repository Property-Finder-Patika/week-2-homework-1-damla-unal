package main

import "fmt"

func main() {
	fmt.Println("---------------Built-in Function New----------------")
	// another way to create a variable is to use the built-in function new
	p := new(int) //p, of type *int, points to an unnamed int variable
	fmt.Printf("%d %d\n", p, *p)

	q := new(int)
	fmt.Printf("%d %d\n", q, *q)

	fmt.Printf("%t\n", p == q)   //false
	fmt.Printf("%t\n", *p == *q) //true

	fmt.Println("---------------Pointer----------------")
	i := 5
	g := &i                       //g points to address of i
	fmt.Printf("%d %d \n", i, *g) //5 5
	fmt.Printf("%d %d \n\n", g, &i)

	i = 6
	fmt.Printf("%d %d \n", i, *g) // 6 6
	fmt.Printf("%d %d \n\n", g, &i)

	*g = 7
	fmt.Printf("%d %d \n", i, *g) //7 7
	fmt.Printf("%d %d \n\n", g, &i)

	fmt.Println("After creating one more pointer.")
	f := g                               // f and g point to address of i
	fmt.Printf("%d %d %d \n", i, *g, *f) // 7 7 7
	fmt.Printf("%d %d %d\n\n", g, f, &i)

	g = increment(i)
	fmt.Printf("%d %d \n", i, *g)   //7 8
	fmt.Printf("%d %d \n\n", g, &i) // different address

	fmt.Println("After setting to nil.")
	g = nil
	fmt.Printf("%d %d \n", i, g) //7 0
	fmt.Printf("%d %d \n", g, &i)

}

func increment(i int) *int {
	fmt.Println("in f()")
	i++
	return &i
}
