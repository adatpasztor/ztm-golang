package main

import "fmt"

func addslice(elem string, slice []string) {
	slice = append(slice, "volt mar itt")
	printslice("kozben", slice)
}

//ez a jo
func addslice2(elem string, slice *[]string) {
	*slice = append(*slice, elem)
	printslice("kozben pointer", *slice)
}

func printslice(title string, slice []string) {
	fmt.Println("------ ", title, "-------")
	for i := 0; i < len(slice); i++ {
		element := slice[i]
		fmt.Println(element)
	}
}

func main() {

	sl1 := []string{"elso", "masodik", "harmadik"}

	printslice("elotte", sl1)
	addslice("meg egy", sl1)
	printslice("utana", sl1)

	fmt.Println("-------------------------------------------------")

	printslice("elotte", sl1)
	addslice2("meg egy ide is", &sl1)
	printslice("utana", sl1)

}
