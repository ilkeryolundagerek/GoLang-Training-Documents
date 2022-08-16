package main

import (
	"fmt"
	"time"
)

func main() {
	//Birinci Örnek: Temel Kanal Yapısı
	transit := make(chan string)
	go func() {
		fmt.Println("Burası haritalama (mapping) işlemleri için.")
		time.Sleep(time.Microsecond * 1000)
		transit <- "haritalamaya uygun 3 kaynak bulundu"
	}()
	gelen := <-transit
	fmt.Println("Harilama goroutine'inden gelen bilgi:", gelen)
	close(transit)

	//İkinci Örnek: Ses iletim fonksiyonları örneği
	mesaj_transit := make(chan string)
	go dinleyici(mesaj_transit)
	giris := <-mesaj_transit
	fmt.Println("Dinleyiciden gelen cevap:", giris)
	close(mesaj_transit)

	//Üçüncü Örnek:
	esyalar := make(chan string, 5)
	esyalar <- "radyo"
	esyalar <- "televiyon"
	esyalar <- "bilgisayar"
	esyalar <- "süpürge"
	esyalar <- "telefon"
	for i := 0; i < 5; i++ {
		go func(esya chan string) {
			e := <-esya
			fmt.Println(e, "eşyası işlenecek.")
		}(esyalar)
	}
	close(esyalar)

	mik := make(chan string, 2)
	mik <- "Merhaba"
	mik <- "Dünya"

	for i := 0; i < 2; i++ {
		sesKarti(mik, 10, 1500)
	}
	close(mik)
	sesKanali := make(chan string, 1)
	mikrofon(sesKanali, "Merhaba dünya")
	sesKarti(sesKanali, 10, 1500)

	var enter string
	fmt.Scanln(&enter)

}

func isci(islem_sonucu chan bool) {
	fmt.Println("Bir sürü iş yapıldı.")
	time.Sleep(time.Microsecond * 2000)
	fmt.Println("İşler bitti")
	islem_sonucu <- true
}

func sesKarti(ses <-chan string, sesSeviyesi int, sure int) {
	fmt.Printf("Ses:\t%s\nSeviye:\t%d\nSüre:\t%d\n", <-ses, sesSeviyesi, sure)
}

func mikrofon(ses chan<- string, mesaj string) {
	ses <- mesaj
}

func dinleyici(mesaj chan string) {
	mesaj <- "Merhaba"
	go func(kanal chan string) {
		kanal <- "Dünya"
	}(mesaj)
	cikis := <-mesaj
	fmt.Printf("\nİç fonksiyon dedi ki: %s\n", cikis)
}
