package main

import (
	"fmt"
	"time"
)

func main() {
	makeChannelPing := make(chan int)
	makeChnnerlPong := make(chan int)

	// go ini adalah goroutine (go)
	go pinger(makeChannelPing, makeChnnerlPong)
	go ponger(makeChannelPing, makeChnnerlPong)

	makeChannelPing <- 1

	for {
		time.Sleep(time.Second)
	}

}

func pinger(pinger <-chan int, ponger chan<- int) {
	for {
		<-pinger
		fmt.Println("ping")
		time.Sleep(time.Second)
		ponger <- 1
	}

}

// tanda panah di sebelah kiri chan (ponger <-chan int) berarti sebagai wadah penerima
// cara baca : variable ponger sebagai wadah penerima nilai integer dari channel

// tanda panah di sebelah kanan chan (pinger chan<- int) berarti sebagai wadah pengirim
// cara baca : variable pinger sebagai wadah pengirim dari nilai integer ke channel
func ponger(pinger chan<- int, ponger <-chan int) {
	for {
		<-ponger
		fmt.Println("pong")
		time.Sleep(time.Second)
		pinger <- 1

	}
}
