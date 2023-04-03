//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

const (
	Active   bool = true
	Inactive bool = false
)

type SecurityTag bool

type Item struct {
	name string
	tag  SecurityTag
}

func activate(tag *SecurityTag) {
	*tag = SecurityTag(Active)
}
func deactivate(tag *SecurityTag) {
	*tag = SecurityTag(Inactive)
}

func checkout(items []Item) {
	for i := 0; i < len(items); i++ {
		deactivate(&items[i].tag)
	}
}

func displayTags(tag *[]Item) {
	for _, t := range *tag {
		fmt.Println("Tag->", t.name, " status->", t.tag)
	}

}

func main() {
	shirt := Item{"Shirt", SecurityTag(Active)}
	pants := Item{"Pants", SecurityTag(Active)}
	purse := Item{"Purse", SecurityTag(Active)}
	watch := Item{"Watch", SecurityTag(Active)}
	items := []Item{shirt, pants, purse, watch}

	fmt.Println("elotte-----------------------------------")
	displayTags(&items)

	fmt.Println("kozben-----------------------------------")
	checkout(items)
	displayTags(&items)

	deactivate(&items[0].tag)

	fmt.Println("utana-----------------------------------")
	displayTags(&items)
	//tags[2].name = "elso"
	//displayTags(&tags)

}
