package main

import "fmt"

func main() {
	number := 3

	if number == 3 {
		fmt.Println("flag-1")
		func() {
			defer fmt.Println("flag 2")
			fmt.Println("flag 3")
		}()
		defer fmt.Println("flag 5")
	}
	fmt.Println("flag4")
}
