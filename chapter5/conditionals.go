package main

import "fmt"

func main() {
	for j := 11; j <= 20; j++ {
		// if - else if - else
		if j % 2 == 0 {
			fmt.Println(j, "even")
		} else {
			fmt.Println(j, "odd")
		}
	}

	for k := 1; k <= 5; k++ {
		// Switch statement
		switch k {
		case 0: fmt.Println("Zero")
		case 1: fmt.Println("One")
		case 2: fmt.Println("Two")
		case 3: fmt.Println("Three")
		case 4: fmt.Println("Four")
		default: fmt.Println("Don't know :(")
		}
	}
}
