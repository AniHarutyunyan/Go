package metanit

import (
	"fmt"
	"testing"
)

func TestPointer(t *testing.T){
	var a int = 4
	var p *int = &a
	fmt.Printf("Address %v \n",p)
	fmt.Printf("Value %v \n",*p)

	*p = 123
	fmt.Printf("Changed [a] value %v %v \n",a,*p)

	f := 2.3
	pf :=&f
	fmt.Printf("value f: %v \n",*pf)
	fmt.Printf("Address f:%v \n",pf)

	var ab *float64
	if ab !=nil{
		fmt.Println(*ab)
	}

	k := new(int)
	fmt.Printf("Value: %v \n",*k)
	*k =15
	fmt.Printf("Value: %v \n",*k)
}

func TestPointer2(t *testing.T){
	d := 5
	fmt.Println("Value before: ",d)
	changeValue(&d)
	fmt.Println("Value after: ",d)

	p1 := createPointer(15)
	fmt.Println("p1:", *p1)
	p2 := createPointer(10)
	fmt.Println("p2: ",*p2)

}

func changeValue(x *int){
	*x = (*x)*(*x)
}

func createPointer(x int)*int{
	p := new(int)
	*p = x
	return p
}