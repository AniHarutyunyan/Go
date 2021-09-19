package metanit

import (
	"fmt"
	"testing"
)

type person struct{
	name string
	age int
	contactInfo contact
}

type contact struct{
	email string
	phone string

}

type node struct{
	value int
	next *node
}
func TestNestedStructs(t *testing.T){

	var tom = person{
		name: "Levon",
		age: 25,
		contactInfo: contact{
			"levon@gm.com",
			"123456",
		},
	}

	fmt.Println(tom.contactInfo.email)
}

func printNodeValue(n *node){
	fmt.Println(n.value)
	if n.next != nil{
		printNodeValue(n.next)
	}
}
func TestNestedStructs2(t *testing.T){

	first := node{value: 5}
	second := node{value: 10}
	third := node{value: 50}

	first.next = &second
	second.next = &third

	var current *node = &first
	for current !=nil{
		fmt.Println(current.value)
		current = current.next
	}

	printNodeValue(&first)

}