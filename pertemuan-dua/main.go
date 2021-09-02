package main

import "fmt"

func main() {
	//cara deklarasi map[tipedata key]tipedata value
	aMapString := map[string]string{
		"fruits": "mangga",
	}
	aMapSlice := map[string][]string{
		"fruits":  {"mangga", "jambu"},
		"animals": {"gajah", "kodok", "ikan"},
	}
	fmt.Println(aMapString)
	fmt.Println(aMapSlice)
}
