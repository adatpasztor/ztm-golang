//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

type Motorcycle struct {
	vehicle_type string
	model_name   string
}
type Car struct {
	vehicle_type string
	model_name   string
}
type Truck struct {
	vehicle_type string
	model_name   string
}

type Lifter interface {
	Lift()
}

func (v Motorcycle) Lift() {
	fmt.Println("Small Lift lifted ", v.model_name)
}
func (v Car) Lift() {
	fmt.Println("Standard Lift lifted ", v.model_name)
}
func (v Truck) Lift() {
	fmt.Println("Large Lift lifted ", v.model_name)
}

func lift_vehicle(t Lifter) {
	t.Lift()
}

func main() {

	m1 := Motorcycle{"Motorcycle", "kek motor"}
	c1 := Car{"Car", "fekete auto"}
	t1 := Truck{"Truck", "zold teherauto"}

	lift_vehicle(c1)

	vehicles := []Lifter{m1, c1, t1}

	for _, v := range vehicles {
		lift_vehicle(v)
	}
}
