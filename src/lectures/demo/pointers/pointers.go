package main

import "fmt"

type Rec struct {
	mezo1 string
	mezo2 string
	szam  int
}

func typeIn(r *Rec, m1 string, m2 string, sz1 int) {
	r.mezo1 = m1
	r.mezo2 = m2
	r.szam = sz1
}

func strReplace(src *string, tgt string) {
	*src = tgt
}

func main() {
	rekord := Rec{}
	szoveg := "Welcome"

	fmt.Println("elotte ->", rekord)
	fmt.Println("elotte ->", szoveg)

	typeIn(&rekord, "siker", "ez is", 1337)
	strReplace(&szoveg, "Siker ez is....")

	fmt.Println("utana ->", rekord)
	fmt.Println("utana ->", szoveg)

}
