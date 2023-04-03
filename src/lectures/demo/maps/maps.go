package main

import "fmt"

func main() {

	var map1 map[string]string
	map1 = make(map[string]string)
	//map2 := make(map)

	map1["PSX"] = "ok"
	map1["PSX2"] = "szinten ok"
	fmt.Println("map1-> ", map1)

}
