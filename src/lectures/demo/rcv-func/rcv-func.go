package main

import "fmt"

type Coordinate struct {
	X, Y int
}

type Mas struct {
	A, B int
	X, Y int
}

func (c *Coordinate) shiftBy(x, y int) {
	c.X += x
	c.Y += y
}
func (m *Mas) shiftBy(x, y int) {
	m.X *= x
	m.Y *= y
	m.A = 1337
	m.B = 88
}

func main() {

	c1 := Coordinate{3, 3}
	var m1 = Mas{1, 2, 4, 4}

	fmt.Println("Elotte c1-> ", c1)
	c1.shiftBy(2, 2)
	fmt.Println("Utana c1-> ", c1)

	fmt.Println("Elotte m1-> ", m1)
	m1.shiftBy(2, 2)
	fmt.Println("Utana m1-> ", m1)

}
