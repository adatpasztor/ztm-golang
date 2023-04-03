package main

import "fmt"

type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}

func occupySpace(lot *ParkingLot, spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) occupySpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) vacateSpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false
}

func (lot *ParkingLot) displayLot() {

	for i, num := range lot.spaces {
		if num.occupied {
			fmt.Println(i, ". parking place is occupied")
		} else {
			fmt.Println(i, ". parking place is free")
		}
		//fmt.Println("parking ", i, ": num-> ", num.occupied)
		//fmt.Printf("type num-> %T\n", num)
	}
}

func main() {
	lot := ParkingLot{spaces: make([]Space, 12)}

	lot.occupySpace(1)
	lot.occupySpace(3)
	lot.occupySpace(8)
	lot.occupySpace(10)
	lot.occupySpace(11)
	lot.occupySpace(12)

	lot.displayLot()

	lot.vacateSpace(12)
	lot.vacateSpace(11)

	lot.displayLot()
}
