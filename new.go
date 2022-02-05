package main

import (
	"fmt"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(1)
	var pesan = make(chan string)
	var katakanHalo = func(siapa string) {
		var data = fmt.Sprintf("hello %s", siapa)
		pesan <- data
	}

	go katakanHalo("bangsat")
	go katakanHalo("kampret")
	go katakanHalo("berak")

	var pesan1 = <-pesan
	fmt.Println(pesan1)

	var pesan2 = <-pesan
	fmt.Println(pesan2)

	var pesan3 = <-pesan
	fmt.Println(pesan3)

}
