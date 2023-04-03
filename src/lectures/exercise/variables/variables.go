//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func t1(a, b int) int {
	return a + b
}

func main() {
	var color = "green"
	var byear, age uint = 1968, 33
	var (
		first_initial string = "elso"
		last_initial  string = "utolso"
	)
	var ageInDays uint
	ageInDays = age * 365

	fmt.Println("color->", color)
	fmt.Println("byear->", byear, " age->", age)
	fmt.Println("first_initial->", first_initial)
	fmt.Println("last_initial->", last_initial)
	fmt.Println("ageInDays->", ageInDays)
	fmt.Println("color->", color)
	fmt.Println("color->", color)

	var result int16 = int16(t1(111, 222))

	fmt.Println("result->", result)

}
