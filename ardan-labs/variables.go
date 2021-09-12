package ardanlabs

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T){
	Variables()
}

func Variables(){
	var a int64
	var b string
	var c float64
	var d bool

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b,b)
	fmt.Printf("var c float64 \t %T [%v]\n",c,c)
	fmt.Printf("var d bool \t %T [%v]\n",d,d )

	aa := 10
	bb := "Ani"
	cc := 15.7
	dd := true

	fmt.Printf("aa := 10 \t %T [%v]\n",aa,aa)
	fmt.Printf("bb := \"Ani\" \t %T [%v]\n",bb,bb)
	fmt.Printf("cc := 15.7 \t %T [%v]\n",cc,cc)
	fmt.Printf("dd := true \t %T [%v]\n",dd,dd)

	aaa := int32(10)
	fmt.Printf("aaa := int32(10) %T [%v]\n",aaa,aaa)
}