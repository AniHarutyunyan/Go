package metanit

import (
	"fmt"
	"testing"
)

func TestMain5(t *testing.T) {
	b := 2 << 2
	c := 16 >> 3

	fmt.Println(b , c)

	a := 5 | 2
	var l int = 6 & 2
	var m int = 5 ^ 2
	var d int = 5 &^ 6

	fmt.Println(a, l,m,d)

}