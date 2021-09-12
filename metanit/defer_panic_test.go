package metanit

import (
	"fmt"
	"testing"
)

func TestDeferPanic(t *testing.T) {
defer finish()
fmt.Println("Program has been started")
fmt.Println("Program is working")

defer fmt.Println("Second defer")

fmt.Println(devide(15,5))
fmt.Println(devide(4,0))
}

func finish(){
	fmt.Println("Program has been finished")
}

func devide(x,y float64)float64{
	if y==0{
		panic("Division by zero!")
	}
	return x / y
}