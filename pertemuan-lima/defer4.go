package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("finish")
	panic("error")
	fmt.Println("start")
}
