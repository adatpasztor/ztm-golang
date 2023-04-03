package main

import "fmt"

func main() {
	a1 := [5]int{11, 22, 33, 44, 55}
	var a2 [5]int
	a3 := [...]int{1, 3, 4}

	s3 := a1[2:5]

	s4 := []int{1, 3}

	s4 = append(s4, 122)

	fmt.Println("a1->", a1)
	fmt.Printf("a1 type-> %T\n", a1)
	fmt.Println("-----------------------------")

	//a2[3] = 23
	fmt.Println("a2->", a2)
	fmt.Printf("a2 type-> %T\n", a2)
	fmt.Println("-----------------------------")

	fmt.Println("a3->", a3)
	fmt.Printf("a3 type-> %T\n", a3)
	fmt.Println("-----------------------------")

	fmt.Println("s3->", s3)
	fmt.Printf("s3 type-> %T\n", s3)
	fmt.Println("-----------------------------")

	fmt.Println("s4->", s4)
	fmt.Printf("s4 type-> %T\n", s4)
	fmt.Println("-----------------------------")

}
