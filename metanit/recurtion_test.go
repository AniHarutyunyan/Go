package metanit

import (
	"fmt"
	"testing"
)

func funtorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * funtorial(n-1)
}

func Fibonacci(n int)int{
	if n==0{
		return 0
	}
	if n==1{
		return 1
	}
	return Fibonacci(n-1)+Fibonacci(n-2)
}

var test = []struct {
	n int
	expected int
}{
	{0, 0},
	{1,1},
    {2,1},
	{3,2},
    {4,3},
    {5,5},
	{6,8},
	{7,11},
}

func TestRecursion(t *testing.T) {
	fmt.Println(funtorial(5))
	fmt.Println(funtorial(4))
	fmt.Println(funtorial(6))
}

func TestFibonacci(t *testing.T){
	for _,tt := range test{
		if Fibonacci(tt.n)!=tt.expected{
			t.Errorf("Expected: %v found %v",Fibonacci(tt.n),tt.expected)
		}
	}
}
