package main

import "fmt"

// Can't be done with type inference.
var t string = "Out of scope variable"

// Don't need to specify type
var (
	a = 10
	b = "B"
	c = -1.5
)

func main() {
	var x string = "Hello World"
	fmt.Println(x)

	var y string
	y = "Hello Again"
	fmt.Println(y)

	z := "Hello for the final time"
	fmt.Println(z)

	// Using the variable from the top level scope
	fmt.Println(t)

	other_func()

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func other_func() {
	fmt.Println(t, "from the other function")
}
