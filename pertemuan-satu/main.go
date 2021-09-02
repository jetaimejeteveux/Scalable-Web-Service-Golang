package main

import (
	"fmt"
)

//assign const dapat dideklarasi di level package
const (
	PHI    = 3.14
	BAHASA = "INDONESIA"
)

//membuat function di package level menggunakan var

func main() {
	fmt.Println("Hello World")
	fmt.Print("Hello World 2")
	//variabel dengan var <namavar> tipe = <value>
	var name string = "Firman"
	fmt.Println(name)
	//variabel <nama var> = <value>
	var tanggal = 24
	fmt.Println(tanggal)

	//<nama var> := <value>
	umur := 24

	fmt.Println(umur)
	//variabel dengan var <namavar> tipe = <value>
	//multivalue
	var kota, universitas string = "yogyakarta", "UPNYK"
	fmt.Println(kota, universitas)

	//<nama var> := <value>
	//multivalue
	firstName, lastName, bulan := "Firman", "Aulia", 3
	fmt.Println(firstName, lastName, bulan)

	//variable menggunakan underscore
	_, lastNameOrang, bulanOrang := "Zidane", "Ibrahim", 5
	fmt.Println(lastNameOrang, bulanOrang)

	//variabel dengan string
	// pointer
	hacktiv := new(string)
	fmt.Print(hacktiv) //akan print address

	//akan error karena alamat jadi tidak bisa diassign value
	//hacktiv = "halo ini coba"

	//untuk isi value harus dereference dengan asterisk
	*hacktiv = "halo ini coba"
	fmt.Println(*hacktiv)

	//ketika mau menghasilkan zero value menggunakna var
	// var zeroInt int akan muncul 0
	//var zeroStr str akan muncul ""

	// var zeroB bool akan muncul false

	//mencoba %s
	var formatInt int = 100
	var formatString string = "string"
	var formatFloat float32 = 3.14
	var formatBool bool = true

	fmt.Printf("format integer %v, format string %v, format float %1.f, format bool %t", formatInt, formatString, formatFloat, formatBool)

	//tipe data
	//tipe data ada uint, int, int32, int64 - pemilihan berdasarkan range value agar dapat mengehmat space di server

	const PHI float32 = 3.14
	fmt.Println(PHI)

	//operator aritmatika
	var add int = 10 + 5
	var min int = 10 - 5
	var multiply int = 5 * 5
	var div int = 10 / 5
	var modulus int = 10 % 10
	var a = 10
	var b = 30
	fmt.Print(add, min, multiply, div, modulus)
	fmt.Printf("hasil tambah %v \n", (a + b))
	fmt.Printf("hasil tambah %v \n", (a - b))
	fmt.Printf("hasil tambah %v \n", (a * b))
	fmt.Printf("hasil tambah %v \n", (a / b))

	//operator perbandingan
	fmt.Println("==", a == b)
	fmt.Println("<", a < b)
	fmt.Println(">", a > b)
	fmt.Println("<=", a <= b)
	fmt.Println(">=", a >= b)
	fmt.Println("!=", a != b)

	//operator logika
	kondisiTrue := true
	kondisiFalse := false
	fmt.Println(kondisiTrue || kondisiFalse)
	fmt.Println(kondisiTrue && kondisiFalse)

}
