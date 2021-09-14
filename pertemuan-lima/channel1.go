package main

import (
	"fmt"
)

func main() {
	//unbuffered channel
	ch := make(chan string)
	printMessage := func(name string) {
		data := fmt.Sprintf("hello %s", name)
		ch <- data // write data to channel
	}
	//syarat untuk menjalankan unbuffered
	//channel harus ada data yang diconsume dan ada 2 goroutine

	go printMessage("andika")
	go printMessage("lia")
	go printMessage("ali")
	// time.Sleep(1 * time.Second)
	message := <-ch //read data from channel
	fmt.Println(message)

	message1 := <-ch //read data from channel
	fmt.Println(message1)

	message2 := <-ch //read data from channel
	fmt.Println(message2)

	// operasi ini akan deadlock
	//message3 := <-ch //read data from channel
	// fmt.Println(message3)
	// a := <- ch berarti membaca value dari ch lalu menempatkan data dari ch ke a
	// ch <- b berarti menulis data dari b ke ch
}
