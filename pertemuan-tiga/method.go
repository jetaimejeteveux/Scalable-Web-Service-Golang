package main

import (
	"fmt"
	"strings"
)

//method harus selalu ada structnya
type person struct {
	name string
	age  int
}

func (p person) sayHello() {
	fmt.Println("hello", p.name)
}

func (p person) getNameAt(i int) string {
	return strings.Split(p.name, " ")[i-1]
}

//method baiknya pake pointer tapi pake seperti di atas(pake return value) ya nggapapa

func (p person) changeName(name string) {
	fmt.Println("change person name value")
	p.name = name
}

func (p *person) changeNamePointer(name string) {
	fmt.Println("change person name value")
	p.name = name
}

func main() {
	var person1 = person{name: "ali baba cc", age: 17}
	person1.sayHello()

	fmt.Println(person1.getNameAt(2))
	fmt.Println("==============================")
	fmt.Println(person1.name)
	person1.changeName("lia")
	fmt.Println("==============================")
	person1.changeNamePointer("lia")
	fmt.Println(person1.name)

}
