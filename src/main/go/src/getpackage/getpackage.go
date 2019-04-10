package main

import (
	"fmt"
	"getpackage/people"
	"getpackage/name"
)

const Country = "Korea"

type mobile struct {
	number int
	deviceId int
}

type car struct {
	name string
	number int
}

func (car) run(){
	fmt.Println("driving")
}

func (mobile) run(){
	fmt.Println("calling")
}

func run(){
	fmt.Println("not working");
}

func main (){
	fmt.Println(Country)

	fmt.Println(name.ReturnMyId())
	fmt.Println(name.AppendCharacter())

	people.SetPeople("ohdoking", 30)
	fmt.Println(people.GetPeople().Name)

	rMobile := mobile{number: 1,deviceId: 1}
	rMobile.run()

	rCar := car{name: "BMW",number: 1}
	rCar.run()

	run()
}