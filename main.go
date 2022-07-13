// this comment
package main

import (
	"fmt"
)

func main() {

	var student1 string
	student1 = "John"

	// type string
	var student2 = "Jane"

	// type is inferred
	x := 2

	var a, b = 6, "Hello"
	c, d := 7, "world"

	fmt.Println(a, b, c, d)
	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

}
