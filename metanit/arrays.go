package main

import "fmt"

func main1() {
	var numbers [10]int = [10]int{0, 10, 6, 5, 5, 5, 5, 5, 5, 5}

	nums := [5]int{1, 5}

	arr := [...]int{4, 5, 5, 5}

	fmt.Println(numbers,nums,arr)

	numbers[3]=15
	fmt.Println(numbers[5],numbers[3])

	array := [3]int{2:15,0:5,1:8}

	fmt.Println(array)
}