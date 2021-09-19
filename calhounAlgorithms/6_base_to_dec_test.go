package calhounAlgorithms

import (
	"fmt"
	"testing"
)

var bases = []struct {
	expected int
	base     int
	n        string
}{
	{15, 2, "1111"},
	{16, 3, "121"},
	{15, 4, "33"},
	{15, 5, "30"},
	{15, 6, "23"},
	{15, 7, "21"},
	{15, 8, "17"},
	{14, 9, "15"},
	{18, 9, "20"},
	{15, 10, "15"},
	{700, 16, "2BC"},
}

func TestBaseToDec(t *testing.T) {
	for _, tt := range bases {
		if BaseToDec(tt.n, tt.base) != tt.expected {
			t.Errorf("Expected: %v found: %v", tt.expected, BaseToDec(tt.n, tt.base))
		}
	}
}

func TestBaseToDec1(t *testing.T) {
	for i, tt := range bases {
		if BaseToDecCharset(tt.n, tt.base) != tt.expected {
			fmt.Println(i + 1)
			t.Errorf("Expected: %v found: %v", tt.expected, BaseToDecCharset(tt.n, tt.base))
		}

	}
}

func TestBaseToBase(t *testing.T) {
	fmt.Println(BaseToBase("1C8", 16, 10))
	fmt.Println(BaseToBase("E", 16, 2))
}
