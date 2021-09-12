package calhounAlgorithms

import "testing"

var tests = []struct {
	list []int
	expectedSum  int
}{
	{[]int{5, 4, 3, 7, 5, 5, 5, 2}, 36},
	{[]int{5, 4, 3, -7}, 5},
	{[]int{}, 0},
	{[]int{0,0,0,0}, 0},
}

func TestSumOfNumbersInList(t *testing.T) {
	for _, tt := range tests {
		sum := SumListOfNumbers(tt.list)
		if sum != tt.expectedSum {
			t.Errorf("Expected %v, found %v", sum, tt.expectedSum)
		}
	}
}

func TestSumOfNumbersInListRec(t *testing.T) {
	for _, tt := range tests {
		sum := SumListOfNumbersRec(tt.list)
		if sum != tt.expectedSum {
			t.Errorf("Expected %v, found %v", sum, tt.expectedSum)
		}
	}
}