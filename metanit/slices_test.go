package metanit

import (
	"fmt"
	"testing"
)

func TestSlices(t *testing.T){
	//var users []string
	//var users []string = []string{"","",""}
	var us = []string{"Ani","Anna","Alla"}
	users := []string{"A","B","C"}

	fmt.Println(us[1])
	users[0]="K"
	fmt.Println(users[0])

	for i,v := range users{
		fmt.Printf("users[%v] = %v\n",i,v)
	}

	var persons []string = make([]string, 5)
	persons = append(persons, "Auuuu")

	for i,v := range persons{
		fmt.Printf("persons[%v] = %v\n",i,v)
	}

	initialUsers := [8]string{"1", "2", "3", "4", "5", "6", "7", "8"}
	users1 := initialUsers[2:5]
	users2 := initialUsers[:4]
	users3 := initialUsers[3:]

	fmt.Println(users1)     
	fmt.Println(users2)     
	fmt.Println(users3)

	//delete element from slice
	initialU := []string{"1", "2", "3", "4", "5", "6", "7", "8"}
	var a int  = 3 //the third one

	initialU = append(initialU[:a],initialU[a+1:]...)

	fmt.Println(initialU)
}