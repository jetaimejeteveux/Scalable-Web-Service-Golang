package main

import "fmt"

func main() {
	//select untuk menspesifikan channel mana yang mau digunakan supaya ngga ketuker

	numbers := []int{5, 2, 3, 7, 4, 8, 9, 1, 6}
	maxCh := make(chan int)
	avgCh := make(chan float64)

	go getMax(numbers, maxCh)
	go getAvg(numbers, avgCh)

	for i := 0; i < 2; i++ {
		select {
		case max := <-maxCh:
			fmt.Println("max numbners : ", max)
		case avg := <-avgCh:
			fmt.Println("avg numbers :", avg)
		}
	}

}

func getMax(numbers []int, ch chan int) {
	maxNum := numbers[0]

	for _, v := range numbers {
		if maxNum < v {
			maxNum = v
		}
	}
	ch <- maxNum
}
func getAvg(numbers []int, ch chan float64) {
	total := 0

	for _, v := range numbers {
		total += v
	}

	avg := float64(total) / float64(len(numbers))
	ch <- avg
}
