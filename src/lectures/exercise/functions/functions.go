//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func greetPerson(name string) string {
	fmt.Println("------------------------------------------")
	fmt.Println("Hi, ", name)
	ret := "Hi, " + name
	return ret
}

func addThree(a, b, c int) int {
	return a + b + c
}

func getTwo() (int, int) {
	return 13, 37
}

func main() {
	greetPerson("PSX")
	fmt.Println("Az uzenet-> ", greetPerson("psx2"))

	fmt.Println("addThree-> ", addThree(1, 2, 333))
	a, b := getTwo()
	fmt.Println("a, b ->  ", a, b)
	// ez nem megy
	//fmt.Println("a, b ->  ", getTwo())

}
