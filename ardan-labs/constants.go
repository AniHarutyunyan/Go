package ardanlabs

import (
	"fmt"
	"testing"
)

func TestD(t *testing.T){
	const ui = 123456
	const uf = 3.14

	const ti int = 123456
	const pi float64 = 3.14

	const(
		A1 = iota
		B1 = iota
		C1 = iota
	)

	fmt.Println("1:",A1,B1,C1)

	const(
		A2 = iota
		B2
		C2
	)

	fmt.Println("2:",A2,B2,C2)

	const(
		A3 = iota+1
		B3
		C3
	)

	fmt.Println("3:",A3,B3,C3)

	const(
		Ldate = 1<<iota
		Ltime
		Lmicrosenconds
		Llongfile
		Lshortfile
		LUTC
	)

	fmt.Println("Log:",Ldate,Ltime,Lmicrosenconds,Llongfile,Lshortfile,LUTC)

	//const a int = 45646546546546546546546546546
	//const b = 45646546546546546546546546546
	//fmt.Println(b)
}