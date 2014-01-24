package main

import "fmt"

func main() {
	i := 1

	// While loop variation
	for i <= 10 {
		fmt.Println(i)
		i += 1
	}

	// Traditional for loop
	for j := 11; j <= 20; j++ {
		fmt.Println(j)
	}
}
