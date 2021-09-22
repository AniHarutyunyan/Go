package calhounAlgorithms

import "fmt"

func FindTwoThatSum(numbers []int, sum int) {
	for i := 0; i < len(numbers); i++ {
		second := sum - numbers[i]
		for j := 0; j < len(numbers); j++ {
			if j == i{
				break
			}
			if second == numbers[j] {
				fmt.Printf("[%v]+[%v] = [%v]\n",numbers[i],numbers[j],sum)
				break
			}else{
				fmt.Println("There is no numbers which sum is ",sum)
			}
		}
	}
}