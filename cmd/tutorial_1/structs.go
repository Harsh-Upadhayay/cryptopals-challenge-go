package main

import (
	"fmt"
)

type gasEngine struct {
	mpg      int
	gallons  int
	engOwner owner
}

type electricEngine struct {
	mpkwh int
	kwh   int
}

type owner struct {
	name string
}

type engine interface {
	milesLeft() int
}

// Method definition.
func (e gasEngine) milesLeft() int {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() int {
	return e.mpkwh * e.kwh
}

// Methods can't be defined on interfaces.
// Therefore this is a generic method.
func canMakeIt(e engine, miles int) bool {
	return e.milesLeft() >= miles
}

func main5() {

	var myEngine gasEngine = gasEngine{25, 15, owner{"Harsh"}}
	myEngine2 := gasEngine{50, 30, owner{"Harsh"}}
	var myEngine3 gasEngine
	myEngine3.mpg = 100
	myEngine3.gallons = 60

	fmt.Printf("Engine 1: %v, %v, %v\nEngine 2: %v, %v, %v, %v, %v\n, Engine 3: %v, %v, %v\n",
		myEngine.mpg, myEngine.gallons, myEngine.milesLeft(),
		myEngine2.mpg, myEngine2.gallons, myEngine2.milesLeft(), myEngine2.engOwner.name, canMakeIt(myEngine2, 50),
		myEngine3.mpg, myEngine3.gallons, myEngine3.milesLeft())
}
