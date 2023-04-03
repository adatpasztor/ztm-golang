package main

import "fmt"

type st1 struct {
	szam   int
	szoveg string
}

func main() {
	var rec1 st1

	rec1.szam = 1337
	rec1.szoveg = "PSX"

	fmt.Println(rec1)
	fmt.Println("struct_proba-> ok")
}
