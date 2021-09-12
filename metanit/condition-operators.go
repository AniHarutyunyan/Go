package main

import "fmt"

func main() {
	a := 15
	b := 15

	if a > b {
		fmt.Printf("a: %v is bigger than b: %v\n",a,b)
	}else if a < b{ 
		fmt.Printf("a: %v is less than b: %v\n",a,b)
	}else{
		fmt.Printf("a: %v is equal to b: %v\n",a,b)
	}

	d := 2
	switch(d+1){
	case 5:
		fmt.Println("d + 1 = ", d +1)
	case 4:
		fmt.Println("d + 1 = ", d + 1)
	case 1, 3, 6:
		fmt.Println("Multiple case")
	default:
		fmt.Println("the variable value is not defined")
	}

}