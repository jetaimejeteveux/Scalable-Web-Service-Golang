package main

import "fmt"

func main() {
	ch := make(chan string)

	names := []string{"lia", "ila", "ali"}

	for _, v := range names {
		go func(name string) {
			data := fmt.Sprintf("hello %s", name)
			ch <- data // write data to channel
		}(v)
	}

	// go func(name string) {
	// 	data := fmt.Sprintf("hello %s", name)
	// 	ch <- data // write data to channel
	// }("ali")
	// printMessage(ch)

	for i := 0; i < 3; i++ {
		printMessage(ch)
	}
	// go func(name string) {
	// 	data := fmt.Sprintf("hello %s", name)
	// 	ch <- data // write data to channel
	// }("lia")
	// printMessage(ch)

	//channel bisa dilempar sebagai parameter
}

func printMessage(ch <-chan string) {
	//function khusus baca atau read
	fmt.Println(<-ch)
}
