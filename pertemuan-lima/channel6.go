package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)
	go sendData(ch)
	retrieveData(ch)

}

func sendData(ch chan<- int) {
	for i := 0; true; i++ {
		ch <- i
		execTime := rand.Int()%10 + 1
		fmt.Println("time exec : ", execTime)
		time.Sleep(time.Duration(rand.Int()%10+1) * time.Second)
	}
}

func retrieveData(ch <-chan int) {
loop:
	for {
		select {
		case data := <-ch:
			fmt.Println(data, "diterima")
		case <-time.After(5 * time.Second):
			fmt.Println("timeout")
			break loop
		}
	}

}
