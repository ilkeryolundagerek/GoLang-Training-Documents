package main

import (
	"fmt"
	"time"
)

func main() {
	go rutin1(0, 150)
	go rutin2()
	go rutin1(-100, -1)

	var enter string
	fmt.Println("\nÇıkış için bir tuça basınız.")
	fmt.Scanln(&enter)
}

func rutin1(min, maks int) {
	for i := min; i < maks; i++ {
		fmt.Print(i, " ")
		time.Sleep(time.Microsecond * 100000)
	}
}

func rutin2() {
	cumle := "There is no one who loves pain itself, who seeks after it and wants to have it, simply because it is pain..."
	for _, h := range cumle {
		fmt.Print(string(h), " ")
		time.Sleep(time.Microsecond * 130500)
	}
}
