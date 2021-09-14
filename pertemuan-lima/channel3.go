package main

import (
	"fmt"
	"time"
)

func main() {
	//buffered channel ngga blocking karena asynchronous
	//unbuffered channel blocking

	//buffered channel digunakan ketika kita sudah tau jumlah data yang akan diproses

	ch := make(chan string, 2)
	//berarti bisa menampung 2 data sekaligus tanpa blocking
	//buffer itu satu goroutine cukup (goroutine main)
	// ch <- "hello"
	// ch <- "hello2"
	// fmt.Println(<-ch)
	// ch <- "hello3"
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	go func() {
		for {
			i := <-ch
			fmt.Println("name", i)
		}
	}()

	// ch <- "hello1"

	for i := 0; i < 100; i++ {
		data := fmt.Sprintf("hello %v", i)
		ch <- data
	}
	time.Sleep(1 * time.Second)
}
