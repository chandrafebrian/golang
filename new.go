package main

import (
	"fmt"
	"time"
)

func main() {
	makeChannel := make(chan string)

	go pengirim(makeChannel)
	go penerima(makeChannel)

	in := 0
	fmt.Scanln(&in)
}

func pengirim(makeChannel chan string) {
	for i := 0; i < 10; i++ {
		makeChannel <- "ping"
	}

}

func penerima(makeChannel chan string) {
	for {
		msg := <-makeChannel
		fmt.Println(msg)
		time.Sleep(time.Second)
	}
}
