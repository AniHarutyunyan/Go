package calhounAlgorithms

import "fmt"

func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("Fizz Buzz")
		} else if i%3 == 0 {
			fmt.Print("Fizz")
		} else if i%5 == 0 {
			fmt.Print("Buzz")
		} else {
			fmt.Print(i, "")
		}
		if i!=n {
			fmt.Print(", ")
		}
	}
}

func FizzBuzzSwitch(n int) {
	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Print("Fizz Buzz, ")
		case i%3 == 0:
			fmt.Print("Fizz, ")
		case i%5 == 0:
			fmt.Print("Buzz, ")
		default:
			fmt.Print(i, ", ")
		}
	}
}
