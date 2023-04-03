//--Summary:
//  Create a program that can perform dice rolls using various configurations
//  of number of dice, number of rolls, and number of sides on the dice.
//  The program should report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must use variables to configure:
//  - number of times to roll the dice
//  - number of dice used in the rolls
//  - number of sides of all the dice (6-sided, 10-sided, etc determined
//    with a variable). All dice must use the same variable for number
//    of sides, so they all have the same number of sides.
//
//--Notes:
//* Use packages from the standard library to complete the project
//* Try using different values for your variables to check the results

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func roll(diceNum int, diceSide int) int {
	ret := rand.Intn(diceSide) + 1
	fmt.Println(diceNum, " kocka-> ", ret)
	return ret
}

func main() {
	const rollNums = 2
	const diceNumber = 2
	const diceSide = 6

	//rand.Seed(45)
	rand.Seed(time.Now().UnixNano())

	var rollSum int

	fmt.Println(rollNums, "dobas", diceNumber, "db", diceSide, "oldalu kockaval")

	rollSum = 0
	for r := 1; r <= rollNums; r++ {
		for d := 1; d <= diceNumber; d++ {
			rollSum += roll(d, diceSide)
		}
	}
	//fmt.Println("rollSum->", rollSum)
	switch {
	case rollSum == 2 && rollNums == 2:
		fmt.Print(rollSum)
		fmt.Println(" Snake eyes !!!")
	case rollSum == 7:
		fmt.Print(rollSum)
		fmt.Println(" Lucky 7 !!!")
	case rollSum%2 == 0:
		fmt.Print(rollSum)
		fmt.Println(" Even !!!")
	default:
		fmt.Print(rollSum)
		fmt.Println(" Odd !!!")

	}
	//	fmt.Println("PSX-> ok")
}
