package main

import (
	"errors"
	"fmt"
	"strings"
)

func validation(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("input jangan kosong brader")
	}
	return true, nil
}
func catch() {
	if r := recover(); r != nil {
		fmt.Println("terjadi kesalahan", r)
	}
}
func main() {
	//membuat custom error
	input := ""

	// _, err := validation(input)
	// if err != nil {
	// 	fmt.Println("input harus string number")
	// 	fmt.Println("error", err)
	// } else {
	// 	fmt.Println("hallo ", input)
	// 	fmt.Println("error", err)
	// }

	if valid, err := validation(input); valid {
		fmt.Println("hallo ", input)
	} else {
		// fmt.Println("error", err.Error())
		panic(err.Error())
	}

	// if err != nil {
	// 	fmt.Println("input harus string number")
	// 	fmt.Println("error", err)
	// } else {
	// 	// fmt.Println("hallo ", input)
	// 	// fmt.Println("error", err)
	// 	panic(err.Error())
	// }

}
