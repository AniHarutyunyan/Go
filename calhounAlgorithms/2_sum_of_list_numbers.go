package calhounAlgorithms

import "fmt"

func SumListOfNumbers(list []int) (sum int) {
	for _, v := range list {
		sum += v
	}
	return
}

func SumListOfNumbersRec(list []int) (sum int) {
	fmt.Println(list)
	if len(list) == 0 {
		return 0
	}
	return list[0] + SumListOfNumbersRec(list[1:])
}