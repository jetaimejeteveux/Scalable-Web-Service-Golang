package main

import "fmt"

func main() {
	//closure pertama
	getMinMax := func(n []int) (int, int) {
		min, max := n[0], n[0]
		for _, v := range n {
			if v < min {
				min = v
			} else if v > max {
				max = v
			}
		}
		return min, max
	}
	numbers := []int{2, 5, 2, 6, 7, 3, 2}
	min, max := getMinMax(numbers)
	fmt.Println(min, max)

	//closure kedua
	newNumber := func(n int) []int {
		var result []int
		for _, v := range numbers {
			if v > n {
				continue
			} 
			result = append(result, v)
		}
		return result
	}
	fmt.Println(newNumber(5))

	banyak, isi := filterMax(numbers, 5)
}

func filterMax (n []int, max int) (int, func() []int){
	var result []int
		for _, v := range n {
			if v > max {
				continue
			} 
			result = append(result, v)
		}
		return len(result), func() []int  {
			return result
			
		}
}

func filter(data []string, callback func(string) bool) []string {
	var result []string

	for _, v := range data {
		if filtered := callback(v); filtered{
			result = append(result, v)
		}
	}
}