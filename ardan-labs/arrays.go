package ardanlabs

import (
	"fmt"
	"testing"
)

func TestE(t *testing.T){
	friends := [5]string{"Ani", "Levon", "Anna", "Ashot", "Alla"}
	fmt.Printf("Bfr[%s]",friends[1])

	for i := range friends{
		friends[i]="Jack"

		if i==1{
			fmt.Printf("Aft[%s]\n",friends[1])
		}
	}

	for i,v := range friends{
		friends[1]="Jack"
		if i==1{
			fmt.Printf("v %s \n",v)
		}
	}

}