package main

import "fmt"

func main() {
	var numberA int = 4
	var newA int = numberA
	var numberB *int = &numberA

	fmt.Println("numberA value :", numberA)
	fmt.Println("number A address:", &numberA)
	fmt.Println("new A value :", newA)
	fmt.Println("numberB value :", *numberB)
	fmt.Println("number B address :", numberB)

}
