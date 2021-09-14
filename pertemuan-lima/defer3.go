package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("finish")
	os.Exit(1)
	fmt.Println("start")
}
