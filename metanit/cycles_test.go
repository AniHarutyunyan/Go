package metanit

import (
	"fmt"
	"testing"
)

func TestMain3(t *testing.T) {
	// for i := 1; i < 10; i++ {
	// 	fmt.Println(i*i)
	// }

	// for i := 0; i < 20; i=i+2 {
	// 	fmt.Println(i)
	// }

	// for i := 1; i < 10; i++ {
	// 	for j := 1; j < 10; j++ {
	// 		fmt.Print(i*j,"\t")
	// 	}
	// 	fmt.Println()
	// }

	array := [4]int{1,2,3,4}
	for a,b := range array{
		fmt.Println("array[",a,"]:",b)
	}

	numbers := [...]int{1,2,3}

	for _,b := range numbers{
		fmt.Println(b)
	}

	nums := [...]int{1,-5,-7,8,9,-9,-9}
	sum := 0
	for _,b :=range nums{
		if b < 0 {
			continue
		}
		if b == 8 {
			break
		}
		sum=sum+b
	}
	fmt.Println(sum)
}