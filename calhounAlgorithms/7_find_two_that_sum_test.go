package calhounAlgorithms

import (
	"fmt"
	"testing"
)

func TestFindTwoThatSum(t *testing.T) {
	numbers := []int{3, 2, 4, 5, 1,6,0,6,8}
	FindTwoThatSum(numbers, 7)
	fmt.Println("-------------------------")
	numbers1 := []int{-3, 2}
	FindTwoThatSum(numbers1, 7)
}
