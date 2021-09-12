package calhounAlgorithms

import (
	"fmt"
	"testing"
)

func MissingNumber(numbers []int)int{
	l :=len(numbers)
	return (l*(l+1)/2)-SumListOfNumbers(numbers)
}
func TestMissingNumber(t *testing.T){
	numbers := []int{0, 2, 3, 4}
	fmt.Printf("Is missing %v\n",MissingNumber(numbers))
}