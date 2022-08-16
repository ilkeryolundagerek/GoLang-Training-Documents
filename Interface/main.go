package main

import "fmt"

func main() {
	vosvos := Araba{id: 1, sofor: "ilker", sahip: "koç holding", ibre: 110}
	jet := Ucak{id: 10, sahip: "Doğan Holding", pilot: "ahmet", yukseklik: 100000}

	araciCalistir(vosvos)
	araciCalistir(jet)

	hepsi := []AracKontrat{vosvos, jet}

	for _, arac := range hepsi {
		araciCalistir(arac)
	}
}

type AracKontrat interface {
	motoruBaslat()
}

func araciCalistir(arac AracKontrat) {
	arac.motoruBaslat()
}

type Araba struct {
	id    int
	sofor string
	sahip string
	ibre  int
}

func (a Araba) motoruBaslat() {
	fmt.Printf("%s motoru çalıştırıp, sürüşe hazılanıyor.\n", a.sofor)
}

type Ucak struct {
	id        int
	sahip     string
	pilot     string
	yukseklik int
}

func (u Ucak) motoruBaslat() {
	fmt.Printf("%s motoru çalıştırıp, uçuşa hazılanıyor.\n", u.pilot)
}
