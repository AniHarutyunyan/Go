package metanit

import (
	"fmt"
	"testing"
)

type Vehicle interface {
	move()
}
type Car struct{}
type Aircraft struct{}

func drive(vehicle Vehicle){
	vehicle.move()
}
func (c Car) move() {
	fmt.Println("Car is running")
}
func (a Aircraft) move() {
	fmt.Println("Aircraft is flying")
}
func TestInterfaces(t *testing.T) {
	var tesla Vehicle = Car{}
	var boing Vehicle = Aircraft{}
	tesla.move()
	boing.move()
	drive(tesla)
	drive(boing)
}
