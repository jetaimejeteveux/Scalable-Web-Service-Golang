package main

import (
	"fmt"
	"reflect"
)

type personLagi struct {
	name    string
	address string
	age     int
}

func (p *personLagi) ChangeName(newName string) {
	p.name = newName
}

func (p *personLagi) getProperty() {
	var reflectVal = reflect.ValueOf(p)

	fmt.Println(reflectVal)
	if reflectVal.Kind() == reflect.Ptr {
		reflectVal = reflectVal.Elem()
	}

	fmt.Println(reflectVal.NumField())

	//mengecek tiap property, apa tipe datanya, apa valuenya
	reflectType := reflectVal.Type()
	for i := 0; i < reflectVal.NumField(); i++ {
		fmt.Println("property name : ", reflectType.Field(i).Name)
		fmt.Println("property type: ", reflectType.Field(i).Type)
		fmt.Println("property type : ", reflectType.Field(i))
		fmt.Println("property value : ", reflectVal.Field(i))
		fmt.Println("============================================")
	}
}

func main() {
	person1 := &personLagi{name: "Ali", address: "jakarta"}
	person1.getProperty()
	//cara gampang person1.changeName("lia")

	reflectVal := reflect.ValueOf(person1)
	methodCN := reflectVal.MethodByName("ChangeName")
	methodCN.Call([]reflect.Value{
		reflect.ValueOf("lia"),
	})
}
