package calhounAlgorithms

import (
	"testing"
)

var words = []struct {
	word         string
	reversedWord string
}{
	{"hello", "olleh"},
	{"world", "dlrow"},
	{"hello bro", "orb olleh"},
	{"", ""},
	{"Reverse a string", "gnirts a esreveR"},
}

var words1 = []struct {
	word         string
	reversedWord string
}{
	{"hello", "olleh"},
	{"寿福", "福寿"},
}

func TestReverseString(t *testing.T){
	for _,tt := range words{
		if ReverseWord(tt.word)!=tt.reversedWord{
			t.Errorf("Expected: %v found: %v",ReverseWord(tt.word),tt.reversedWord)
		}
	}
}

func TestReverseWordRune(t *testing.T){
 for _,tt := range words1{
	if ReverseWordRune(tt.word)!=tt.reversedWord{
		t.Errorf("Expected: %v found: %v",ReverseWordRune(tt.word),tt.reversedWord)
	}
 }
}

func TestReverseStringBuilder(t *testing.T){
	for _,tt := range words{
		if ReverseStringBuild(tt.word)!= tt.reversedWord{
			t.Errorf("Expected: %v found: %v",ReverseStringBuild(tt.word),tt.reversedWord)
		}
	}

}