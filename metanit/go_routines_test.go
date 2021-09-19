package metanit

import (
	"fmt"
	"testing"
)

func TestPoli(t *testing.T) {

	for i := 0; i < 7; i++ {
		go factorial(i)
	}
	fmt.Println("Enter something")
	fmt.Scanln()
	fmt.Println("End")
}

func TestGoRout(t *testing.T) {
	for i := 0; i < 7; i++ {
		go func(n int) {
			result := 1
			for j := 1; j <= n; j++ {
				result *= j
			}
			fmt.Println("Factorial for: ", n, " is ", result)
		}(i)
	}
	fmt.Scanln()
	fmt.Println("End")
}

func factorial(n int) {
	if n == 0 {
		fmt.Println("Factorial for: ", n, " is ", 1)
		return
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	fmt.Println("Factorial for: ", n, " is ", result)
}
