package main

import (
	"fmt"
	"reflect"
)

func main() {
	//reflect berguna untuk menginspeksi variabel

	var message string = "hello world"

	reflectVal := reflect.ValueOf(message)
	fmt.Println(reflectVal)
	fmt.Println("tipe data", reflectVal.Type())
	if reflectVal.Kind() == reflect.String {
		fmt.Println("value", reflectVal.String())
	}

	//cara membedakan kind dan type
	//bandingkan struct
	//best practice untuk mengecek value adalah dengan menggunakan kind, karena menggunakan typeof mengembalikan interfce jadi tidak bisa dicek
	//reflect biasa digunakan untuk komparasi tipe data
	//karena biasanya untuk membandingkan tipe data ngga perlu value langsung ke tipe data aja yang dipake

}
