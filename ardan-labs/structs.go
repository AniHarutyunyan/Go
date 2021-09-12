package main

import "fmt"

type example struct{
	counter int64
	pi float32
	flag bool
}
func main(){
	var e1 example
	fmt.Printf("%+v\n",e1)

	e2 := example{
		flag: true,
		pi: 3.14,
		counter: 10,
	}

	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)

	e3 := struct{
		flag bool
		counter int16
	}{
		flag: true,
		counter: 10,
	}

	fmt.Println(e3)

	bill := struct{
		flag bool
		counter int16
	}{
		flag: true,
		counter: 10,
	}

	nancy := struct{
		flag bool
		counter int16
	}{
		flag: true,
		counter: 10,
	}

	//bill = nancy
	fmt.Println(bill == nancy)
}