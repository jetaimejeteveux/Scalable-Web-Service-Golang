package main

import (
	"fmt"
	"runtime"
	"sync"
)

func printNumber(mode string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Println(mode, i)
	}
	wg.Done()
}

func main() {
	// printNumber("biasa")
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	var wg sync.WaitGroup
	wg.Add(1)
	go printNumber("goroutine", &wg)
	wg.Wait()
}
