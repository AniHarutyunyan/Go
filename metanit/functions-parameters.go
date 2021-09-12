package main

import "fmt"

func main4() {
	a := 20

	fmt.Println(Foo(a))

	fmt.Println(Add(47,10))

	fmt.Println(Add1(1,5,7,10,0))

	fmt.Println(Add1([]int{1,5,7,10,1}...))

    sum, diff := Add3(14,21)
	fmt.Println(sum,diff)

}

func Foo(a int) string {
	if a == 10 {
		return "Ani"
	}
	return "Levon"
}

func Add(x, y int)int{
return x+y
}

func Add1(numbers ...int)int{
	sum := 0
	for _,v := range numbers {
		sum = sum+v
	}
	return sum
}

func Add2(numbers ...int)(sum int){
	for _,v := range numbers {
		sum = sum+v
	}
	return
}

func Add3(x,y int)(sum, diff int){
	return x+y,x-y
}
