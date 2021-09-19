package metanit

import (
	"fmt"
	"testing"
)


type user struct{
	name string
	age int
}
func TestStructs(t *testing.T){
	var u user = user{"Ani",  15}
	u1 := user{name:"Levon",age: 20}
	u3 := user{}

	fmt.Println(u.age,u1.name,u3.age)

	tom := user{name: "Tom",age: 22}
	var tomPointer *user =&tom
	tomPointer.age = 14
	fmt.Println(tom.age)
	(*tomPointer).age = 23
	fmt.Println(tom.age)

	var anna *user = &user{"Anna",25}
	var bob *user = new(user)
	
	fmt.Println(anna.age)
	fmt.Println(bob.name)
	
	alla := user{"Karen",23}
	var agePointer *int = &alla.age
	*agePointer = 26
	fmt.Println(alla.age)


}