package main

import "fmt"

func main() {
	ch := make(chan int)
	go writeMessage(ch)
	readMessage(ch)
}

func writeMessage(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		if i == 5 {
			close(ch)
			break
		}
	}
	close(ch)
}

func readMessage(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
		//rangenya channel cuma satu aja balikannya, ngga ada indexnya
	}
}
