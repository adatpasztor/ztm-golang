package main

import "fmt"

type proba1 struct {
	szam    int
	szoveg  string
	szoveg2 string
}

type proba_interfesz interface {
	kiir()
	kiir2()
}

func (p proba1) kiir() {
	fmt.Println("szam->", p.szam, " szoveg->", p.szoveg, "szoveg2->", p.szoveg2)
}
func (p proba1) kiir2() {
	fmt.Println("----------")
}

func pr(a proba_interfesz) {
	a.kiir()
}

//-----------------------------------------------------------
// pointer tipussal

type pp_interfesz interface {
	pp_kiir()
	pp_modosit(int)
}

func (p *proba1) pp_kiir() {
	fmt.Println("szam->", p.szam, " szoveg->", p.szoveg, "szoveg2->", p.szoveg2)
}

func (p *proba1) pp_modosit(v int) {
	p.szam = v
}

func pp_pr(a pp_interfesz) {
	a.pp_kiir()
}

func main() {
	var c proba_interfesz
	c = proba1{33, "harminharom", "semmi"}

	a := proba1{szam: 12, szoveg: "tizenketto", szoveg2: "semmi"}
	//b := proba1{13, "tizenharom", "ez is semmi"}
	//a.kiir()
	//b.kiir()
	//b.kiir2()
	pr(a)

	pr(c)
	pr(proba1{22, "huszonketto", "-"})

	fmt.Println("-------- pelda a pointer tipusu metodusokra -----------------")
	// mind a ketto jo
	//var pp pp_interfesz
	//pp = &proba1{55, "otvenot", "shemmi"}
	pp := &proba1{55, "otvenot", "shemmi"}
	pp.pp_kiir()
	pp.pp_modosit(133)
	pp.pp_kiir()

}
