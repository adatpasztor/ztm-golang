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

type Status bool

const (
	active   Status = true
	inactive Status = false
)

type Tag struct {
	name   string
	status Status
}

func setStatus(tag *Tag, status Status) {
	tag.status = status
}

func checkout(tag *[]Tag) {
	for _, t := range *tag {
		t.status = inactive
	}
}

func displayTags(tag *[]Tag) {
	for _, t := range *tag {
		fmt.Println("Tag->", t.name, " status->", t.status)
	}

}

func main() {
	tags := []Tag{
		{"egy", active},
		{"ketto", active},
		{"harom", active},
		{"negy", active},
	}

	displayTags(&tags)

	setStatus(&tags[1], inactive)

	displayTags(&tags)
	tags[2].name = "elso"
	displayTags(&tags)

}
