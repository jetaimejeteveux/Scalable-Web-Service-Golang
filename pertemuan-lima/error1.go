package main

import (
	"fmt"
	"strconv"
)

func main() {
	//kalo kita terjun ke golang
	//maka akan selalu dan pasti untuk periksa error

	//error ini interface jadi atiati
	strNum := "1"

	num, err := strconv.Atoi(strNum)

	if err != nil {
		fmt.Println("input harus string number")
		fmt.Println("error", err)
	} else {
		fmt.Println("format benar : ", num)
		fmt.Println("error", err)
	}
}
