package metanit

import (
	"fmt"
	"testing"
)

type mile int
//type kilometer int

func distanceToEnemy (distance mile){
	fmt.Println("расстояние для противника:")
	fmt.Println(distance,"mile")
}

type library []string

func printBooks(lib library){
	for _,val := range lib{
		fmt.Println(val)
	}
}
func TestTypes(t *testing.T){

	var distance mile =21
	fmt.Println(distance)
	distance+=5
	fmt.Println(distance)

	var distan mile = 5
	distanceToEnemy(distan)
	// var distance2 kilometer = 5
	// distanceToEnemy(kilometer2)

	var myLib library = []string{"B1","B2","B3"}
	printBooks(myLib)
}