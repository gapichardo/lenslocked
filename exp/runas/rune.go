package main

import (
	"fmt"
)

func main() {
	const r1 = '€'
	fmt.Println("(int32) r1:", r1)
	fmt.Printf("(HEX) r1: %x\n", r1)
	fmt.Printf("(as a String) r1: %s\n", r1)
	fmt.Printf("(as a character) r1: %c\n", r1)

	// Here, you see that a byte slice is a collection of runes and that printing
	// a byte slice with fmt.Println() might not return what you expected.
	// In order to convert a rune into a character, you should use %c in a
	// fmt.Printf() statement. In order to print a byte slice as a string,
	// you will need to use fmt.Printf() with %s.

	fmt.Println("A string is a collection of runes:", []byte("Mihalis"))
	aString := []byte("Mihalis")
	for x, y := range aString {
		fmt.Println(x, y)
		fmt.Printf("Char: %c\n", aString[x])
	}
	fmt.Printf("%s\n", aString)
}
