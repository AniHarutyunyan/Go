package metanit

import (
	"fmt"
	"strconv"
	"testing"
)

func Add4(x,y int) int{
	return x+y
}

func Multiply(x,y int)int{return x*y}
func DisplayMessage(message string){fmt.Println(message)}
func Operation(x,y int, operation func(int, int) int)int{
	result := operation(x,y)
	return result
}

func isEven(n int)bool{return n%2==0}
func isPositive(n int)bool{return n>0}

func sum(numbers []int, criteria func (n int) bool)int{
	result:=0

	for _,v := range numbers{
		if criteria(v) {
			result+=v
		}
	}
	return result
}

//
func Substruct(x,y int)int{return x-y}
func selectFn(n int)func(int, int)int{
	if n==1{
		return Add4
	}else if n==2{
		return Substruct
	}else{
		return Multiply
	}
}

func TestB(t *testing.T){
	var f func(int, int) int = Add4
	fmt.Println(f(3,4))

	var x = f(4,5)
	fmt.Println(x)

	multi := Multiply(8,9)
	DisplayMessage(strconv.Itoa(multi))


	DisplayMessage(strconv.Itoa(Operation(4,8,Add4)))
	DisplayMessage(strconv.Itoa(Operation(10,15,Multiply)))


	numbers := []int{1,-5,7,1}
	nums := []int{1,2,4,1}

	DisplayMessage(strconv.Itoa(sum(nums,isEven)))
	DisplayMessage(strconv.Itoa(sum(numbers,isPositive)))

	a := selectFn(1)
	b := selectFn(2)
	c := selectFn(3)

	DisplayMessage(strconv.Itoa(a(4,5)))
	DisplayMessage(strconv.Itoa(b(4,5)))
	DisplayMessage(strconv.Itoa(c(4,5)))
}