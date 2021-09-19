package calhounAlgorithms

import (
	"testing"
)
var nm = []struct {
	n        int
	base        int
	expected string
}{
	{15,2,"1111"},
	{15,3,"120"},
	{15,4,"33"},
	{15,5,"30"},
	{15,6,"23"},
	{15,7,"21"},
	{15,8,"17"},
	{15,9,"16"},
	{15,10,"15"},
	{16,2,"10000"},
	{16,3,"121"},
	{16,4,"100"},
	{16,5,"31"},
	{5,2,"101"},
	{1,2,"1"},
	{7,3,"21"},
	{18,16,"12"},
	{25,16,"19"},
	{7,16,"7"},
	{700,16,"2BC"},
	{16,16,"10"},
	
}

func TestDecToBase(t *testing.T){
	for _,tt := range nm{
		if DecToAnyBase(tt.n,tt.base)!=tt.expected{
			t.Errorf("Expected: %v found: %v",tt.expected,DecToAnyBase(tt.n,tt.base))
		}
	}
}
func TestDecToBaseCharset(t *testing.T){
	for _,tt := range nm{
		if DecToBase(tt.n,tt.base)!=tt.expected{
			t.Errorf("Expected: %v found: %v",tt.expected,DecToBase(tt.n,tt.base))
		}
	}
}