package metanit

import (
	"fmt"
	"testing"
)

func TestChan(t *testing.T){
	intChan := make(chan int)

	go func(){
		fmt.Println("Go routine starts")
		intChan<- 5
	}()
	fmt.Println(<-intChan)
	fmt.Println("End")
}

func TestChannel(t *testing.T){
	intChan := make(chan int)

	go fact(5,intChan)
	fmt.Println(<-intChan)
	fmt.Println("The end")
}

func fact(n int, ch chan int){
	result := 1
	for i := 1; i <= n; i++ {
		result*=i
	}
	fmt.Println("Factorial for: ", n, " is ", result)
	ch <-result
}




