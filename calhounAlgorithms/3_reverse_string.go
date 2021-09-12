package calhounAlgorithms

import (
	"fmt"
	"strings"
)

func ReverseWord(word string) (reversedWord string) {
	wordLen := len(word)
	for i := 1; i <= wordLen; i++ {
		reversedWord += string(word[wordLen-i])
	}
	fmt.Println("Reverted word:" + reversedWord)
	return
}
func ReverseWordRune(word string) (reversedWord string) {
	var res string
	for _,v := range word {
		res = string(v) + res
	}
	fmt.Println("Reverted word:" + res)
	return res
}
func ReverseStringBuild(word string) string {
	wordLen := len(word)
	var sb strings.Builder
	for i := 1; i <= wordLen; i++ {
		sb.WriteByte(word[wordLen-i])
	}
	fmt.Println("Reverted word:" + sb.String())
	return sb.String()
}
