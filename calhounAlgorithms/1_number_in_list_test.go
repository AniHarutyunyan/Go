package calhounAlgorithms

import (
	"fmt"
	"testing"
)

var test = []struct {
	list []int
	num  int
	out  bool
}{
	{[]int{5, 4, 3, 7, 5, 5, 5, 2}, 3, true},
	{[]int{5, 4, 4, 7, 5, 5, 5, 2}, 3, false},
    {nil, 3, false},
    {[]int{5, -4, -3, 7, 5, 5, 5, 2}, 3, false},
    {[]int{}, 3, false},
    {[]int{5, 4, 3, 7, -5, 5, 5, 2}, -5, true},
}

func TestNumberInList(t *testing.T) {
	for _, tt := range test {
        isInList := NumberInList(tt.list, tt.num)
        fmt.Printf("Find %v in %v\n",tt.num,tt.list)
		if isInList != tt.out {
			t.Errorf("was %v, want %v", isInList, tt.out)
		}
	}
}