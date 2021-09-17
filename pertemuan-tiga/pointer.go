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

	number := 10
	fmt.Println(number)

	change(&number, 100)
	fmt.Println(number)

}
func change(number *int, newNumber int) {
	*number = newNumber
}
