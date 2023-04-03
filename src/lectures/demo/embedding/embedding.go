package main

import "fmt"

const (
	Small = iota
	Medium
	Large
)

const (
	Ground = iota
	Air
)

type Beltsize int
type Shipping int

// a const hasznalataval a nevet mint string kapjuk vissza
// gyakorlatilag toString-kent mukodik a String receiver fuggveny
func (b Beltsize) String() string {
	return []string{"Small", "Medium", "Large"}[b]
}

func (s Shipping) String() string {
	return []string{"Ground", "Air"}[s]
}

type Conveyor interface {
	Convey() Beltsize
}

type Shipper interface {
	Ship() Shipping
}

type WarehouseAutomator interface {
	Conveyor
	Shipper
}

type SpamMail struct {
	amount int
}

func (s SpamMail) String() string {
	return "Spam mail"
}

func (s SpamMail) Ship() Shipping {
	return Air
}

func (s SpamMail) Convey() Beltsize {
	return Small
}

func automate(item WarehouseAutomator) {
	fmt.Printf("Convey %v on %v conveyor\n", item, item.Convey())
	fmt.Printf("Ship %v via %v \n", item, item.Ship())

}

type ToxicWaste struct {
	weight int
}

func (t *ToxicWaste) Ship() Shipping {
	return Ground
}

func main() {
	//	var a Beltsize = Small

	// fmt.Println("a->", a)
	// fmt.Printf("a-> %v \n", a)
	// fmt.Println("a->", Beltsize(1))

	mail := SpamMail{40000}
	automate(mail)

	//waste := ToxicWaste{20000}

	//automate(&waste)

}
