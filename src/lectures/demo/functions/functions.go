package main

import "fmt"

func double(a int) int {
	return a + a
}

func add(j, b int) int {
	return j + b
}

func greet() {
	fmt.Println("Hello from greet function")
}

func main() {
	greet()
	dozen := double(6)
	fmt.Println("dozen->", dozen)
	bakersDozen := add(double(6), 1)
	fmt.Println("bakersDozen->", bakersDozen)

}
