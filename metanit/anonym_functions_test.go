package metanit

import (
	"fmt"
	"testing"
)

func selectFunc(n int)(func (int,int )int){
	if n ==1{
		return func (x, y int)int{return x+y}
	}else if n==2{
		return func( x, y int)int{return x*y}
	}else{
		return func(x, y int)int{return x-y}
	}
}

func square() func()int{
	x := 2

	return func()int{
		x++
		return x * x
	}
}

func TestA(t *testing.T){
	f := func(x,y int)int{return x+y}

	fmt.Println(f(10,15))
	fmt.Println(f(14,16))

	foo := selectFunc(1)
	foo1 := selectFunc(2)
	foo2 := selectFunc(3)

	fmt.Println(foo(10,18))
	fmt.Println(foo1(10,18))
	fmt.Println(foo2(10,18))

	foo3 := square()
	fmt.Println(foo3())
	fmt.Println(foo3())
	fmt.Println(foo3())
}