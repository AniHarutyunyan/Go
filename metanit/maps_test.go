package metanit

import (
	"fmt"
	"testing"
)

func TestMaps(t *testing.T){
	var persons = map[string]int{
		"Tom": 1,
		"Bob": 2,
		"Sam": 4,
		"Alice": 8,
	}
	for key,val := range persons{
		fmt.Printf("persons[%v]: %v\n",key,val)
	}

	if val,ok := persons["Sam"]; ok{
		fmt.Println(val)
	}

	persons["Alice"]=123123
	fmt.Println(persons["Alice"])

	persons["Anna"]=456
	persons["Janna"]=789

	for key,val := range persons{
		fmt.Printf("persons[%v]: %v\n",key,val)
	}

	delete(persons,"Alice")

	for key,val := range persons{
		fmt.Printf("persons[%v]: %v\n",key,val)
	}
}