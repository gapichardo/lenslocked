package main

import (
	"fmt"
)

func main() {
	const sLiteral = "\x99\x42\x32\x55\x50\x35\x23\x50\x29\x9c"
	fmt.Println(sLiteral)
	fmt.Printf("x: %x\n", sLiteral)

	fmt.Printf("sLiteral length: %d\n", len(sLiteral))
	for i := 0; i < len(sLiteral); i++ {
		fmt.Printf("%x ", sLiteral[i])
	}
	fmt.Println()

	// Here, you can see that you can access a string literal as if it is a slice. 
	// Using %q in fmt.Printf() with a string argument will print a double-quoted 
	// string that is safely escaped with Go syntax. Using %+q in fmt.Printf() with 
	// a string argument will guarantee ASCII-only output.

	// Last, using % x (note the space between the % character and the x character) 
	// in fmt.Printf() will put spaces between the printed bytes. In order to print 
	// a string literal as a string, you will need to call fmt.Printf() with %s.
	
	fmt.Printf("q: %q\n", sLiteral)
	fmt.Printf("+q: %+q\n", sLiteral)
	fmt.Printf(" x: % x\n", sLiteral)

	// Here you define a string named s2 with three Unicode characters. 
	// Using fmt.Printf() with %#U will print the characters in the U+0058 format. 
	// Using the range keyword on a string that contains Unicode characters will 
	// allow you to process its Unicode characters one by one.

	// The output of len(s2) might surprise you a little. 
	// As the s2 variable contains Unicode characters, 
	// its byte size is larger than the number of characters in it.
	fmt.Printf("s: As a string: %s\n", sLiteral)
	s2 := "€£³"
	for x, y := range s2 {
		fmt.Printf("%#U starts at byte position %d\n", y, x)
	}

	fmt.Printf("s2 length: %d\n", len(s2))
	const s3 = "ab12AB"
	fmt.Println("s3:", s3)
	fmt.Printf("x: % x\n", s3)

	fmt.Printf("s3 length: %d\n", len(s3))

	for i := 0; i < len(s3); i++ {
		fmt.Printf("%x ", s3[i])
	}
	fmt.Println()
}
