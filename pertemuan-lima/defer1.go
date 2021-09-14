package main

import "fmt"

func main() {
	defer fmt.Println("finish")

	fmt.Println("start1")
	fmt.Println("start2")
	fmt.Println("start3")
	// return
	// fmt.Println("start")

	// fmt.Println("start")

	orderSomeFood("nasi goreng")
	orderSomeFood("mie")

	//biasanya defer digunakan untuk close connection
	//setelah query misalnya open connection ke db, lansung defer close connection
}

func orderSomeFood(food string) {
	defer fmt.Println("finish")
	if food == "mie" {
		fmt.Println("--------------")
		fmt.Println("anda memesan", food)
		fmt.Println("--------------")
		return
	}
	fmt.Println("anda memesan", food)
}
