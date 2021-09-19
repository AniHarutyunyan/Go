package calhounAlgorithms

import (
	"fmt"
)

func DecToBase(n, base int) string {
	const charset = "0123456789ABCDEF"
	dec := ""
	for n != 0 {
		m := n % base
		dec = string(charset[m])+dec
		n = n / base
	}
	return dec
}

func DecToAnyBase(n int, base int) string {
	dec := ""
	for n != 0 {
		m := n % base
		dec = fmt.Sprintf("%X%s",m,dec)
		n = n / base
	}
	return dec
}