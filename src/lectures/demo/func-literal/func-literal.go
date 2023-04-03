package main

import "fmt"

func kibocsato(kezdo int) func() int {
	return func() int {
		kezdo++
		return kezdo
	}
}

func add(lhs, rhs int) int {
	return lhs + rhs
}

func compute(lhs, rhs int, op func(a, b int) int) int {
	fmt.Printf("Compute-> %v & %v\n", lhs, rhs)
	return op(lhs, rhs)
}

func main() {

	// tiz := kibocsato(10)
	// szaz := kibocsato(100)

	// fmt.Println("tiz ->", tiz())
	// fmt.Println("tiz ->", tiz())
	// fmt.Println("szaz ->", szaz())
	// fmt.Println("tiz ->", tiz())
	// fmt.Println("szaz ->", szaz())

	//fmt.Println("2 + 2 =", compute(2, 2, add))

	mul := func(a, b int) int {
		fmt.Printf("Szorzas-> %v * %v =", a, b)
		return a * b
	}

	fmt.Println("3 + 5 =", compute(3, 5, add))
	fmt.Println("3 * 5 =", compute(3, 5, mul))

}
