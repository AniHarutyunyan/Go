package calhounAlgorithms

import (
	"fmt"
	"math"
	"strconv"
)

func BaseToDec(base string, toBase int) int {
	sum := 0
	for i := 1; i <= len(base); i++ {
		str, err := strconv.Atoi(string(base[len(base)-i]))
		if err != nil {
			er := fmt.Errorf("cannot convert string to int")
			fmt.Println(er)
		}
		sum = sum + str*int(math.Pow(float64(toBase), float64(i-1)))

	}
	return sum
}

func BaseToDecCharset(base string, toBase int) int {
	charset := "0123456789ABCDEF"
	sum := 0
	multiplier := 0
	for i := 1; i <= len(base); i++ {
		str := base[len(base)-i]
		for j := 1; j < len(charset); j++ {
		
			if charset[j] == str{
				multiplier =j
				break
			}
		}
		sum = sum + multiplier*int(math.Pow(float64(toBase), float64(i-1)))

	}

	return sum
}
func BaseToBase(number string, fromBase int, toBase int) string {
	return DecToBase(BaseToDecCharset(number,fromBase),toBase)
}
