package main

import "fmt"

func hello() {
	// can't access the fmt package without importing it, because they are not shared in the same package
	fmt.Println("Hey, I am Damla!")
	bye() // no need to import a package, because this function is in the package-scope of the main package.
}
