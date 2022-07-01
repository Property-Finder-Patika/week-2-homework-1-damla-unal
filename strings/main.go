package main

import (
	"fmt"
	"strings"
)

func main() {
	// it's a raw string literal
	json := `
{
	"Items": [{
		"Item": {
			"name": "Teddy Bear"
		}
	}]
}`

	fmt.Println(json)
	fmt.Println("********************")

	fmt.Println(strings.ToLower("Damla Unal")) //damla unal
	msg := `
	
	The weather looks good.
I should go and play.
	`
	fmt.Println(msg)
	fmt.Println(strings.TrimSpace(msg))

	fmt.Printf("\n%s\n", "Literals with Underscore")
	fmt.Println(1_000_000_000)
	fmt.Printf("%d\n", 0b1001_1111)
	fmt.Printf("%d\n", 0xFFFF_FFFF)

	// an english letter (search web for: ascii control code)
	var letter byte = 'A'
	fmt.Println("an english letter:", letter)

	// a non-english letter (search web for: unicode codepoint)
	var unicode rune
	unicode = 'Ç'
	fmt.Println("a non-english letter:", unicode)
}
