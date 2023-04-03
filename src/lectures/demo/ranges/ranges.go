package main

import "fmt"

func main() {
	sl1 := []string{"Hi", "PSX", " -> ", "ok."}

	for i, elem := range sl1 {
		fmt.Println("i->", i, "elem->", elem)
		for _, betu := range elem {
			fmt.Printf("%q", betu)
		}
		fmt.Println("")
	}
}
