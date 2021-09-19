package metanit

import (
	"fmt"
	"testing"
)

func TestMethods(t *testing.T) {
	var lib library = library{"B1", "B2", "B3"}
	lib.print()

	var p person = person{
		name: "Levon",
		age:  25,
		contactInfo: contact{
			"levon@gm.com",
			"123456",
		}}

	p.print1()
	p.eat("cake")
	var pp *person = &p
	fmt.Println(p.age)
	pp.updateAge(44)
	fmt.Println(p.age)
	p.updateAge(15)
}

func (l library) print() {
	for _, val := range l {
		fmt.Println(val)
	}
}
func (p person) print1() {
	fmt.Println("Name: ", p.name)
	fmt.Println("Age: ", p.age)
}
func (p person) eat(meal string) {
	fmt.Println(p.name + " is eating " + meal)
}
func(p *person)updateAge(newAge int){
	(*p).age = newAge
}
